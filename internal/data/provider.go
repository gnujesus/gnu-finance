package data

type DataFetcher interface {
	Fetch(q Query) (CompanyInfo, error)
}
