package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alvarezjulia/wine-dal/src/clients"
	"github.com/olivere/elastic/v7"
)

// EntryDao is an interface that wraps the entryDao struct type.
// Purpose: Can be mocked for testing
var (
	EntryDao entryDaoInterface = &entryDao{}
)

type entryDaoInterface interface {
	GetDoc(ctx context.Context, wineID string) (Wine, error)
	CreateNewDoc(ctx context.Context, wine Wine) (string, error)
	GetManyDocs(ctx context.Context, wineID string) (WineList, error)
	CreateManyDocs(ctx context.Context, wines WineList) error
}

type entryDao struct{}

func (e *entryDao) GetDoc(ctx context.Context, wineID string) (Wine, error) {
	query := elastic.NewBoolQuery()
	musts := []elastic.Query{elastic.NewTermQuery("id", wineID)}
	query = query.Must(musts...)
	searchResult, err := clients.ElasticClient.Get().Index("bottles").Id(wineID).Do(ctx)
	if err != nil {
		return Wine{}, err
	}

	var wine Wine
	err = json.Unmarshal(searchResult.Source, &wine)
	if err != nil {
		return Wine{}, err
	}

	return wine, nil
}

func (e *entryDao) CreateNewDoc(ctx context.Context, wine Wine) (string, error) {
	ID := trimSpaces(wine.Title) + trimSpaces(wine.Vintage) + trimSpaces(wine.Winery)
	res, err := clients.ElasticClient.Index().
		Index("bottles").
		BodyJson(wine).
		Id(ID).
		Do(ctx)
	if err != nil {
		return "", err
	}
	fmt.Printf("Indexed wine to index %s with id: %s\n", res.Id, res.Index)
	return res.Id, nil
}

func (e *entryDao) GetManyDocs(ctx context.Context, wineID string) (WineList, error) {
	var wines WineList
	return wines, nil
}

func (e *entryDao) CreateManyDocs(ctx context.Context, wines WineList) error {
	bulk := clients.ElasticClient.Bulk().Index("bottles")
	for _, wineEntry := range wines.WineSlice {
		wine := Wine{
			Country:      wineEntry.Country,
			Region:       wineEntry.Region,
			Lage:         wineEntry.Lage,
			Sweetness:    wineEntry.Sweetness,
			SugarLevel:   wineEntry.SugarLevel,
			WineType:     wineEntry.WineType,
			WineColor:    wineEntry.WineColor,
			Title:        wineEntry.Title,
			Description:  wineEntry.Description,
			AlcoholLevel: wineEntry.AlcoholLevel,
			Vintage:      wineEntry.Vintage,
			ValidEAN:     wineEntry.ValidEAN,
			Acidity:      wineEntry.Acidity,
			Winery:       wineEntry.Winery,
			Grape:        wineEntry.Grape,
			Appellation:  wineEntry.Appellation,
		}
		ID := trimSpaces(wine.Title) + trimSpaces(wine.Vintage) + trimSpaces(wine.Winery)
		bulk.Add(elastic.NewBulkIndexRequest().Id(ID).Doc(wine))
	}

	if _, err := bulk.Do(ctx); err != nil {
		fmt.Println("Failed to bulk load")
		return err
	}
	return nil
}

func trimSpaces(withSpace string) string {
	withoutSpace := strings.Join(strings.Fields(withSpace), "")
	return withoutSpace
}
