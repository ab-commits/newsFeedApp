# newsFeedApp
Building a Contextual News Data Retrieval System
Used Go as the Backend Language for developemnt and Mongo Db for storing the data in form of documents

WorkFlow
1. Run the script inside scripts folder either run go run load_data.go or ./scripts.exe .It will load the MongoDb with the data supplied.
2. Run Http Server(Implemented via Gorrila Mux framework) to access the specified EndPoints as mentioned in the document supplied.
