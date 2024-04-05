# Schedule

- Ngày học lí thuyết : 19/2 -> 23/2 trừ ngày 21/2
- Ngày viết code example : 26/2 -> 27/2
- [Báo cáo](https://docs.google.com/document/d/1gulcQ9RLnbEhr09rynkKHy5FFzhoIlqPYSPKGNDmVuo)

# Description

<p sytle="color='#22d3ee'">Source code cơ bản, nó triển khai CURD API với chỉ mục order trong Elastic-search</p>

Định nghĩa 1 chỉ mục Order như sau:
```
type OrderItemModel struct {
	ProductId int     `json:"product_id"`
	Amount    float64 `json:"amount"`
	Quantity  int     `json:"quantity"`
}

type OrderModel struct {
	PurchaseAt  time.Time        `json:"purchased_at"`
	Lines       []OrderItemModel `json:"lines"`
	TotalAmount float64          `json:"total_amount"`
	SalesMan    struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"salesman"`
	SalesChanel string `json:"sales_chanel"`
	Status      string `json:"status"`
}
```


<p sytle="color='#22d3ee'">Kiến thức Elasticsearch áp dụng được: </p>

-  Thêm/sửa/xóa chỉ mục
-  Search query
-  Aggregations
-  Nested

# Running
Using shell script in /scitp
```
sh ./script/run.sh
```
## Các API được triển khai:
-	 GET    /elastic/order/search-all
-  GET    /elastic/order/search-ids
-  GET    /elastic/order/stats-order
-  POST   /elastic/order/update-order-by-id
-  POST   /elastic/order/create-order 
-  DELETE /elastic/order/delete-order 
