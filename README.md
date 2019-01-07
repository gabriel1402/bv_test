# BeenVerified - DataDeck Engineer Technical Challenge

## Description
API that exposes a data source of songs and genres. Developed with Go and Goji.

## Requirements
- Go
- SQLite
- GNU Compiler Collection (gcc)

## Installation
Clone the project in your Go workspace `src` directory with ```git clone https://github.com/gabriel1402/bv_test.git```

Enter the cloned directory (bv_test) and run ```glide install``` to set the project dependencies. It may be necessary to have a GNU compiler in order to install the SQLite driver for Go.

The main package of the project is `datadeck`, so you can generate the project executable with ```go install .\datadeck\```.

After installing the project, return to your workspace. You can either run the `datadeck.exe` file generated in your workspace's `bin` directory or execute ```go run .\src\bv_text\datadeck\``` to get the server up. This will be running at `localhost:8000`.

## API Reference

### Songs list
Return the list of songs in json format. Searches over artists, song names and genre's names.

* **URL**

  /songs

* **Method:**

  `GET`
  
*  **URL Params**

   **Optional:**
 
   `query=[alphanumeric]`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `[{"ID":2,"Artist":"424","Song":"Gala","Genre":"Indie Rock","Length":189},{"ID":3,"Artist":"Colornoise","Song":"Amalie","Genre":"Noise Rock","Length":246}] | null`

### Songs list (ordered by length)
Return the list of songs in json format, ordered by length.
* **URL**

  /songs/byLength

* **Method:**

  `GET`
  
*  **URL Params**

   **Optional:**
 
   `max=[numeric]`
   `min=[numeric]`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `[{"ID":3,"Artist":"Colornoise","Song":"Amalie","Genre":"Noise Rock","Length":246},{"ID":2,"Artist":"424","Song":"Gala","Genre":"Indie Rock","Length":189}] | null`

### Genres list
Returns the list of genres, along with the amount of songs and total length.
* **URL**

  /genres

* **Method:**

  `GET`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `[{"Name":"Pop","Songs":5,"TotalLength":1001},{"Name":"Rap","Songs":2,"TotalLength":408}]`
