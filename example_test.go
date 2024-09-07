package orderedjson_test

import (
	"fmt"
	orderedjson "orderedJson"
	"strings"
)

const testJSON = `{
  "order": {
    "orderHeader": {
      "externalOrderID": "ACSIT1000064330",
      "orderType": "CHANGE",
      "orderTypeAlias": "Modify",
      "orderContext": "Modify|Add/Remove/Modify VAS",
      "orderChannel": "VR App",
      "orderStatus": "SUBMITTED",
      "orderProcessingStatus": "PENDING",
      "reasonCode": "",
      "reasonDescription": "",
      "customerRequestedDate": "2022-09-18T15:49:30.069+08:00",
      "dealerCode": "BCC10000KUL",
      "salesAgentCode": "1463",
      "partnerID": "",
      "resellerID": "BCC10000KUL",
      "capturedBy": "S0617103764",
      "capturedOn": "2022-09-18T15:49:30.069+08:00",
      "orderParameters": {
        "attribute": [
          {
            "attributeName": "POST_SALE_ENRICHMENT_FLAG",
            "attributeValue": "N"
          },
          {
            "attributeName": "orderContext",
            "attributeValue": "Modify|Add/Remove/Modify VAS"
          },
          {
            "attributeName": "orderChannel",
            "attributeValue": "App"
          },
          {
            "attributeName": "DEALER_NAME",
            "attributeValue": "Dealer Name"
          },
          {
            "attributeName": "REGION_ID",
            "attributeValue": "R01323"
          },
          {
            "attributeName": "UPFRONT_PAYMENT",
            "attributeValue": "Y"
          }
        ]
      }
    },
    "commonInformation": {
      "customer": {
        "customerID": "C47957735"
      },
      "customerAccount": {
        "customerAccountID": "A11293852",
        "customerId": "C47957735"
      }
    },
    "orderLines": [
      {
        "orderLineActionCommand": "ADD_OR_REMOVE_PRODUCTS",
        "orderLineStatus": "SUBMITTED",
        "orderLineProcessingStatus": "SUBMITTED",
        "subscriberOrderItem": {
          "subscriber": {
            "customerAccountID": "A11293852",
            "customerId": "C47957735",
            "subscriberID": "S10342847",
            "subscriberProducts": [
              {
                "productID": "PI100002",
                "productStatus": "I",
                "priceOverride": "N",
                "productOfferingRefID": "NVF00160",
                "linkedProductOfferingID": "PB32172"
              },
              {
                "productID": "PI100005",
                "productStatus": "O",
                "priceOverride": "N",
                "productOfferingRefID": "NVF00190",
                "linkedProductOfferingID": "PB32172"
              }
            ]
          }
        }
      }
    ],
    "payment": [
      {
        "paymentPlace": "Card",
        "paymentItem": [
          {
            "amount": {
              "currency": "GBP",
              "amount": "188.0"
            },
            "paymentMethod": {
              "name": "Card"
            }
          }
        ]
      }
    ],
    "orderNotes": [
      {
        "createdGroupID": "S0617103764",
        "notes": "MODIFY_ORDER_VALIDATION"
      }
    ]
  },
  "interactionID": "ACSIT1000064330",
  "interactionDate": "2022-09-19T15:49:29.984+08:00",
  "sourceApplicationID": "amazon",
  "serviceName": "createOrder",
  "triggeredBy": "S0617103764",
  "lang": "ENG",
  "channel": "CSA",
  "opID": "HOB",
  "buID": "DEFAULT"
}`

// ExampleDecoder_orderedPath shows an error but if you compare the outputs they are identical
func ExampleDecoder_orderedPath() {
	v, err := orderedjson.Unmarshal([]byte(testJSON))
	if err != nil {
		panic(err)
	}
	showV(v, 0)

	// Output:
	//0 order
	//.. 0 orderHeader
	//.... 0 externalOrderID ACSIT1000064330
	//.... 1 orderType CHANGE
	//.... 2 orderTypeAlias Modify
	//.... 3 orderContext Modify|Add/Remove/Modify VAS
	//.... 4 orderChannel VR App
	//.... 5 orderStatus SUBMITTED
	//.... 6 orderProcessingStatus PENDING
	//.... 7 reasonCode
	//.... 8 reasonDescription
	//.... 9 customerRequestedDate 2022-09-18T15:49:30.069+08:00
	//.... 10 dealerCode BCC10000KUL
	//.... 11 salesAgentCode 1463
	//.... 12 partnerID
	//.... 13 resellerID BCC10000KUL
	//.... 14 capturedBy S0617103764
	//.... 15 capturedOn 2022-09-18T15:49:30.069+08:00
	//.... 16 orderParameters
	//...... 0 attribute [
	//........ 0 [
	//.......... 0 attributeName POST_SALE_ENRICHMENT_FLAG
	//.......... 1 attributeValue N
	//........ ]
	//........ 1 [
	//.......... 0 attributeName orderContext
	//.......... 1 attributeValue Modify|Add/Remove/Modify VAS
	//........ ]
	//........ 2 [
	//.......... 0 attributeName orderChannel
	//.......... 1 attributeValue App
	//........ ]
	//........ 3 [
	//.......... 0 attributeName DEALER_NAME
	//.......... 1 attributeValue Dealer Name
	//........ ]
	//........ 4 [
	//.......... 0 attributeName REGION_ID
	//.......... 1 attributeValue R01323
	//........ ]
	//........ 5 [
	//.......... 0 attributeName UPFRONT_PAYMENT
	//.......... 1 attributeValue Y
	//........ ]
	//...... ]
	//.. 1 commonInformation
	//.... 0 customer
	//...... 0 customerID C47957735
	//.... 1 customerAccount
	//...... 0 customerAccountID A11293852
	//...... 1 customerId C47957735
	//.. 2 orderLines [
	//.... 0 [
	//...... 0 orderLineActionCommand ADD_OR_REMOVE_PRODUCTS
	//...... 1 orderLineStatus SUBMITTED
	//...... 2 orderLineProcessingStatus SUBMITTED
	//...... 3 subscriberOrderItem
	//........ 0 subscriber
	//.......... 0 customerAccountID A11293852
	//.......... 1 customerId C47957735
	//.......... 2 subscriberID S10342847
	//.......... 3 subscriberProducts [
	//............ 0 [
	//.............. 0 productID PI100002
	//.............. 1 productStatus I
	//.............. 2 priceOverride N
	//.............. 3 productOfferingRefID NVF00160
	//.............. 4 linkedProductOfferingID PB32172
	//............ ]
	//............ 1 [
	//.............. 0 productID PI100005
	//.............. 1 productStatus O
	//.............. 2 priceOverride N
	//.............. 3 productOfferingRefID NVF00190
	//.............. 4 linkedProductOfferingID PB32172
	//............ ]
	//.......... ]
	//.... ]
	//.. ]
	//.. 3 payment [
	//.... 0 [
	//...... 0 paymentPlace Card
	//...... 1 paymentItem [
	//........ 0 [
	//.......... 0 amount
	//............ 0 currency GBP
	//............ 1 amount 188.0
	//.......... 1 paymentMethod
	//............ 0 name Card
	//........ ]
	//...... ]
	//.... ]
	//.. ]
	//.. 4 orderNotes [
	//.... 0 [
	//...... 0 createdGroupID S0617103764
	//...... 1 notes MODIFY_ORDER_VALIDATION
	//.... ]
	//.. ]
	// 1 interactionID ACSIT1000064330
	// 2 interactionDate 2022-09-19T15:49:29.984+08:00
	// 3 sourceApplicationID amazon
	// 4 serviceName createOrder
	// 5 triggeredBy S0617103764
	// 6 lang ENG
	// 7 channel CSA
	// 8 opID HOB
	// 9 buID DEFAULT
}

func showV(v []*orderedjson.OrderedJson, level int) {
	for i, v := range v {
		switch vt := v.V.(type) {
		case []*orderedjson.OrderedJson: // Next Level Objects
			fmt.Println(strings.Repeat("..", level), i, v.K)
			showV(vt, level+1)
		case []interface{}: // Array of values
			fmt.Println(strings.Repeat("..", level), i, v.K, "[")
			level++
			for j, a := range vt {
				fmt.Println(strings.Repeat("..", level), j, "[")
				switch arrayItem := a.(type) {
				case []*orderedjson.OrderedJson:
					showV(arrayItem, level+1)
				default:
					fmt.Println(strings.Repeat("..", level), j, a)
				}
				fmt.Println(strings.Repeat("..", level), "]")
			}
			level--
			fmt.Println(strings.Repeat("..", level), "]")
		default:
			fmt.Println(strings.Repeat("..", level), i, v.K, v.V)

		}
	}
}
