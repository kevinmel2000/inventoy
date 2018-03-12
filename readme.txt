JSON:

http://localhost:8080/items -> POST
{
  "sku" : "SSI-D00791091-XL-BWH",
  "name" : "Zalekia Plain Casual Blouse (XL,Broken White)",
  "stock" : 137
}

http://localhost:8080/inbound_items -> POST
{
  "itemId" : 1,
  "status" : 1,
  "orderAmount" : 54,
  "receivedAmount" : 54,
  "price" : 77000,
  "total" : 4158000,
  "receiptNumber" : "20170823-75140",
  "notes" : "2017/08/26 terima 54"
}

http://localhost:8080/outbound_items -> POST
{
  "notes" : "Order 1",
  "item" : [
    {
      "itemId" : 1,
      "sellAmount" : 1,
      "price" : 115000,
      "total" : 115000,
      "notes" : "Order 1"    
    },
    {
      "itemId" : 2,
      "sellAmount" : 2,
      "price" : 130000,
      "total" : 130000,
      "notes" : "Order 1"    
    }
  ]
  
}

http://localhost:8080/report/item_value -> GET REPORT ITEM

http://localhost:8080/report/selling/03/2018 -> GET REPORT SELLING
note : 03 -> Month
       2018 -> Year
Make sure you input correct format. 