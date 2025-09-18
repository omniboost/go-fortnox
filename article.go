package fortnox

type ArticleRows []ArticleRow

type ArticleRow struct {
	URL                string `json:"@url,omitempty"`               // Direct URL to the record.
	ArticleNumber      string `json:"ArticleNumber"`                // Unique identifier of the article.
	Description        string `json:"Description"`                  // Description of the article.
	DisposableQuantity string `json:"DisposableQuantity,omitempty"` // The quantity of the article that is available for sale.
	EAN                string `json:"EAN,omitempty"`                // EAN code of the article.
	Housework          bool   `json:"Housework,omitempty"`          // If the article is a housework article.
	PurchasePrice      string `json:"PurchasePrice,omitempty"`      // Purchase price of the article.
	QuantityInStock    string `json:"QuantityInStock,omitempty"`    // The quantity of the article that is currently in stock.
	ReservedQuantity   string `json:"ReservedQuantity,omitempty"`   // The quantity of the article that is reserved.
	SalesPrice         string `json:"SalesPrice,omitempty"`         // Sales price of the article.
	StockPlace         string `json:"StockPlace,omitempty"`         // Stock place of the article.
	StockValue         string `json:"StockValue,omitempty"`         // The total value of the article in stock.
	Unit               string `json:"Unit,omitempty"`               // Unit of the article.
	VAT                string `json:"VAT,omitempty"`                // VAT percentage of the article.
	WebshopArticle     bool   `json:"WebshopArticle,omitempty"`     // If the article is a webshop article.
}

type Article struct {
	URL                       string  `json:"@url,omitempty"`
	Active                    bool    `json:"Active,omitempty"`
	ArticleNumber             string  `json:"ArticleNumber,omitempty"`
	Bulky                     bool    `json:"Bulky,omitempty"`
	ConstructionAccount       int     `json:"ConstructionAccount,omitempty"`
	CostCalculationMethod     string  `json:"CostCalculationMethod,omitempty"`
	DefaultStockLocation      string  `json:"DefaultStockLocation,omitempty"`
	DefaultStockPoint         string  `json:"DefaultStockPoint,omitempty"`
	Depth                     int     `json:"Depth,omitempty"`
	Description               string  `json:"Description,omitempty"`
	DirectCost                float64 `json:"DirectCost,omitempty"`
	DisposableQuantity        float64 `json:"DisposableQuantity,omitempty"`
	EAN                       string  `json:"EAN,omitempty"`
	EUAccount                 int     `json:"EUAccount,omitempty"`
	EUVATAccount              int     `json:"EUVATAccount,omitempty"`
	Expired                   bool    `json:"Expired,omitempty"`
	ExportAccount             int     `json:"ExportAccount,omitempty"`
	FreightCost               float64 `json:"FreightCost,omitempty"`
	Height                    int     `json:"Height,omitempty"`
	Housework                 bool    `json:"Housework,omitempty"`
	HouseworkType             string  `json:"HouseworkType,omitempty"`
	Manufacturer              string  `json:"Manufacturer,omitempty"`
	ManufacturerArticleNumber string  `json:"ManufacturerArticleNumber,omitempty"`
	Note                      string  `json:"Note,omitempty"`
	OtherCost                 float64 `json:"OtherCost,omitempty"`
	PurchaseAccount           int     `json:"PurchaseAccount,omitempty"`
	PurchasePrice             float64 `json:"PurchasePrice,omitempty"`
	QuantityInStock           float64 `json:"QuantityInStock,omitempty"`
	ReservedQuantity          float64 `json:"ReservedQuantity,omitempty"`
	SalesAccount              int     `json:"SalesAccount,omitempty"`
	SalesPrice                float64 `json:"SalesPrice,omitempty"`
	StockAccount              int     `json:"StockAccount,omitempty"`
	StockChangeAccount        int     `json:"StockChangeAccount,omitempty"`
	StockGoods                bool    `json:"StockGoods,omitempty"`
	StockPlace                string  `json:"StockPlace,omitempty"`
	StockValue                float64 `json:"StockValue,omitempty"`
	StockWarning              float64 `json:"StockWarning,omitempty"`
	SupplierName              string  `json:"SupplierName,omitempty"`
	SupplierNumber            string  `json:"SupplierNumber,omitempty"`
	Type                      string  `json:"Type,omitempty"`
	Unit                      string  `json:"Unit,omitempty"`
	VAT                       float64 `json:"VAT,omitempty"`
	WebshopArticle            bool    `json:"WebshopArticle,omitempty"`
	Weight                    int     `json:"Weight,omitempty"`
	Width                     int     `json:"Width,omitempty"`
	CommodityCode             string  `json:"CommodityCode,omitempty"`
}
