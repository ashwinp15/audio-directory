## Overview
This is a graphql API for a server to store short audio bits called ``noobles``. It was a hiring challenge which I had a lot of fun developing
and learned a lot from. <br>
As of now there is only the API implementation which can be tested using the instructions given below. Shortly, I am also planning to adding a frontend to this project.

## Requisites
You need to have the Altair GraphQL client which provides very easy to use and feature rich client (in this case, easy file uploads). <br>
Download it for your system [here](https://altair.sirmuel.design/#download) or install the [chrome-extension](https://chrome.google.com/webstore/detail/altair-graphql-client/flnheeellpciglgpaodhkhmapeljopja?hl=en) or [firefox-extension](https://addons.mozilla.org/en-US/firefox/addon/altair-graphql-client/).

This server is hosted on and EC2 instance. 
All the requests will be served on the following endpoint. Copy this URL into your Altair address bar for making the requests. <br>
http://ec2-3-237-97-25.compute-1.amazonaws.com:8080/query


## Usage 
Below are example requests for all the CRUD operations.

### Query: Read all noobles 
```graphql
{
	 noobles { 
	 #subfields for noobles are necessary
		 id
		 title
		 category 
	 }
}
```
### Query: Read single nooble via id
```graphql
{
  nooble (id: "required_nooble_id") {
	# Again, subfields are required in the query
    title
    id
  }
}
```
The id's are I'm using are auto generated UUIDs from Postgres (overkill for a test project :P) and complicated as such. So an easy way is to issue a "read all noobles" query with id subfield and copy one particular id from the results.

### Mutation: Create user 
```graphql
mutation {
  createUser(input: {
    name: "jon_doe"
    email: "jon_doe@example.com"
    password: "hardtoguess"
  })
  {
	 # response sub fields required 
    email
    name
  }
}
```

### Mutation: Create nooble 

#### Important
In order to create a nooble, the creator(user) must exist beforehand. Queries without the creator field or a non-existent creator will fail. So, make sure to do a create user request first. And then use the email in the creator field of the create nooble request. <br><br>

My plan was to implement authentication and take the creator info via sessions. But cookie handling is not as straightforward in graphQL. I'll need to read up on contexts. For now authentication is not implemented and until I get a hang of using contexts, this is just a constraint I've kept in place.

```graphql
mutation($nooble_file: Upload!) {
	 createNooble(input: {
		 title: "test_title"
		 description: "this is a sample description"
		 category: "people_and_blogs"
		 creator: "jon_doe@example.com"
		 file: $nooble_file
	 })
}
```
#### File uploads
* In the above query, the file to be uploaded is represented by the variable "nooble_file".
* In Altair, towards the bottom of the VARIABLES section, there is an "Add files" option. 
* Click the "Add files" and select the file you want to upload.
* Next to your uploaded file, edit the text field and enter "nooble_file" there (without quotes). This will map the uploaded file to the variable.
* Send the request

### Mutation: Update nooble

The update mutation allows you to change the **title**, **description** and **category** fields of an existing nooble. 

```graphql
mutation {
  updateNooble(id: "required_nooble_id", input: {
    title: "changed_title" 
    category: "changed_title" 
    description: "changed_title" 
  }) 
}
```

### Mutation: Delete nooble

```graphql
mutation {
  deleteNooble(id: "required_nooble_id")
}
```
