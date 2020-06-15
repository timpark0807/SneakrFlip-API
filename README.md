# SneakrFlip API Documentation
The SneakrFlip API is a RESTful service. The API uses standard HTTP response codes, authentication, and verbs. 

The base URL for each request is `https://api.sneakrflip.com/`. 

## Authorization
API requests are authorized using the JWT token obtained from a successful oAuth login. This token must be passed in the header of each request. 

The JWT token will contain information about the current user. The backend will check that the current user matches the user associated with the item in the database. If the users do not match, we deny access to that request. As a result, a User A cannot delete an item created by another User B.

A successful request will return a `200 Successful` response. An unauthorized request will return a `403	Forbidden` response. 

# API Endpoints

## Retrieve an item.
Retrieves an item with a given ID. 

**Endpoint:**

`GET`  `https://api.sneakrflip.com/api/item/:id`

**Parameters:** 
|Name  | Type  | Description |
|--|--|--|
|ID|string|The item's unique identifier.|


## List items for a user
Returns all items created by a user.

**Endpoint:**
`GET` `https://api.sneakrflip.com/api/item`

**Parameters:** 
None 


## Create an item
Creates a new item for a user. 

**Endpoint:**
`POST` `https://api.sneakrflip.com/api/item`

**Parameters:** 
|Name  | Type  | Description |
|--|--|--|
|category|string|Label the item as a Shoe, Clothing, or Other.|
|brand|string|The origin or manufacturer of the item.|
|description|string|A brief description of the item.|
|condition|string|Label the item condition as Deadstock, Very Near Deadstock, or Used.|
|sold|boolean|Whether the item is sold. Default to unsold.|


## Update an item
Updates an existing item in the inventory of a user. Any parameter values not specified will not be modified. 

**Endpoint:**
`PUT` `https://api.sneakrflip.com/api/item`

**Parameters:** 
|Name  | Type  | Description |
|--|--|--|
|category|string|Label the item as a Shoe, Clothing, or Other.|
|brand|string|The origin or manufacturer of the item.|
|description|string|A brief description of the item.|
|condition|string|Label the item condition as Deadstock, Very Near Deadstock, or Used.|
|sold|boolean|Whether the item is sold. Default to unsold.|


## Delete an item
Deletes an item from a user's inventory. 

**Endpoint:**
`DELETE` `https://api.sneakrflip.com/api/item/:id`

**Parameters:** 
|Name  | Type  | Description |
|--|--|--|
|ID|string|The item's unique identifier.|


## Update Status of an Item
Updates the "sold" status of an item based on the id. 

**Endpoint:**
`DELETE` `https://api.sneakrflip.com/api/item/:id`

**Parameters:** 
|Name  | Type  | Description |
|--|--|--|
|ID|string|The item's unique identifier.|



# Objects

## Item
An item in the user's inventory. 

**Attributes:**
|Name  | Type  | Description |
|--|--|--|
|category|string|Label the item as a Shoe, Clothing, or Other.|
|brand|string|The origin or manufacturer of the item.|
|description|string|A brief description of the item.|
|condition|string|Label the item condition as Deadstock, Very Near Deadstock, or Used.|
|sold|boolean|Whether the item is sold. Default to unsold.|
|created by|string|Email address of the user. This field is auto generated.|
|updated on|string|Datetime stamp of the last edit. This field is auto generated. |
