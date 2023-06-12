# dozer-match
 
How to run the system : 

Backend : 

Do `cd backend`

Do `docker compose up`

Go to New terminal and Do `cd backend`

Copy the .env file content and paste it in new terminal


Run the migration file by running `go run main.go migration`

Run `go run main.go server`

Server will start running 


Frontend : 

Open new terminal 

Do `cd dozer-match-frontend `

Do `npm install `

Do `npm run start`


Documentation : 

The dozer match scraper scrapes a site for dozers and returns a list of dozers. Based on that list we create custom filters 
on the frontend which help us to easily navigate through the list.

Backend : 
The backend is responsible for scraping the site cat.com to obtain list of dozers. The backend is designed to extend to multiple sites as well.

Scrapers scrape the page to generate raw html and pass it to modularizers which break the html to chunks each one to be a particular object. Then these html chunks are passed to parsers which convert them to objects. Parsers also filter out
noisy data. 

We use 2 options for parsers normal html parser and gpt parser. The parser is selected during server start when we pass an env
 variable use_gpt (bool) . We also need to pass the gpt key to use gpt parser.

While a scrape of any site is going on we set a redis key to identify that site. Further requests to scrape the site will return the old scrape until the new scrape is completed.

There are get requests to retrive latest scrape.

Scrape data is stored in postgres with a unique scrapeindex . The scrapeindex can be used to identify the scrape. The latest scrape index is also stored in redis and get calls reference it.


Frontend : 

Frontend fetches the latest scrape when its first loaded and displays it along with custom filters pertaining to data generated by the scrape.

We  can request new scrape by hitting rescrape. 







