

delete etalase


https://gql.tokopedia.com/graphql/DeleteShopShowcase



[{"operationName":"DeleteShopShowcase","variables":{"input":{"id":"24412741"}},"query":"mutation DeleteShopShowcase($input: ParamDeleteShopShowcase!) {\n  deleteShopShowcase(input: $input) {\n    success\n    message\n    __typename\n  }\n}\n"}]


res 
[
    {
        "data": {
            "deleteShopShowcase": {
                "success": true,
                "message": "Berhasil menghapus Etalase Toko",
                "__typename": "MutationResult"
            }
        }
    }
]



showcase list

https://gql.tokopedia.com/graphql/ShopShowcase
[{"operationName":"ShopShowcase","variables":{"withDefault":false},"query":"query ShopShowcase($withDefault: Boolean) {\n  shopShowcases(withDefault: $withDefault) {\n    result {\n      id\n      name\n      count\n      uri\n      __typename\n    }\n    error {\n      message\n      __typename\n    }\n    __typename\n  }\n}\n"}]

[
    {
        "data": {
            "shopShowcases": {
                "result": [
                    {
                        "id": "24412741",
                        "name": "Tas Olahraga",
                        "count": 0,
                        "uri": "https://www.tokopedia.com/schmart/etalase/tas-olahraga",
                        "__typename": "showcaseData"
                    },
                    {
                        "id": "24413746",
                        "name": "Tas OutDoor",
                        "count": -4,
                        "uri": "https://www.tokopedia.com/schmart/etalase/tas-outdoor",
                        "__typename": "showcaseData"
                    },
                    {
                        "id": "24627872",
                        "name": "HeadPhone",
                        "count": 0,
                        "uri": "https://www.tokopedia.com/schmart/etalase/headphone",
                        "__typename": "showcaseData"
                    }
                ],
                "error": {
                    "message": "",
                    "__typename": "ErrorResult"
                },
                "__typename": "ShopShowcasesResult"
            }
        }
    }
]


create showcase

https://gql.tokopedia.com/graphql/addShopShowcase

[{"operationName":"addShopShowcase","variables":{"input":{"name":"fhky"}},"query":"mutation addShopShowcase($input: ParamAddShopShowcase!) {\n  addShopShowcase(input: $input) {\n    success\n    message\n    createdId\n    __typename\n  }\n}\n"}]


[
    {
        "data": {
            "addShopShowcase": {
                "success": true,
                "message": "Berhasil menambah Etalase Toko",
                "createdId": "34908887",
                "__typename": "MutationResult"
            }
        }
    }
]