# gutendex-api

This project is a part of assignment 1 from PROG2005 Cloud Technologies, NTNU.
The goal of the assignment is to develop a REST web application that provides data about books in a given language on 
the Gutenberg library. Instead of using the public endpoint, this application uses locally hosted versions. 

## Root 
This end point is under "/librarystats/v1/" and contains links to each endpoint mentioned under. But endpoints that need 
a parameter will not return any data as the parameter needs to be given in the url.

## Bookcount Endpoint

This endpoint returns details about the given language(s).

### Request
```
Method: GET
URL Path: bookcount/?language={:two_letter_language_code+}/
```

### Response
- Content type: ```application/json```
- Status code: 200 if everything is OK, appropriate error code otherwise.
#### Body: example url (bookcount/?language=fi,no/) 
```
[
  {
     "language": "no",
     "books": 21,
     "authors": 14,
     "fraction": 0.0005
  },
  {
     "language": "fi",
     "books": 2798,
     "authors": 228,
     "fraction": 0.0671
  }
]
```

## Readership Endpoint
This endpoint returns the number of potential readers in a given language.

### Request
```
Method: GET
Path: readership/{:two_letter_language_code}{?limit={:number}}
```

### Response
- Content type: application/json
- Status code: 200 if everything is OK, appropriate error code otherwise.

#### Body: example url (readership/no/?limit=2/)
```
[ 
  {
     "country": "Norway",
     "isocode": "NO",
     "books": 21,
     "authors": 14,
     "readership": 5379475
  },
  {
     "country": "Svalbard and Jan Mayen",
     "isocode": "SJ",
     "books": 21,
     "authors": 14,
     "readership": 2562
  }
]
```

## Diagnostics Endpoint

### Request


```
Method: GET
Path: status/
```

### Response
- Content type: application/json
- Status code: 200 if everything is OK, appropriate error code otherwise.

```
{
   "gutendexapi": "<http status code for gutendex API>",
   "languageapi: "<http status code for language2countries API>", 
   "countriesapi": "<http status code for restcountries API>",
   "version": "v1",
   "uptime": <time in seconds from the last service restart>
}
```

## Deployment
The service is deployed on Render under this url: https://assignment-1-q0xw.onrender.com.

## Known issues

- Author count is not 100% accurate. 
