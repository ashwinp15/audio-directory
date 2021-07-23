## Usage

### Create user mutation
```graphql
mutation {
  createUser(input: {
    name: "jon_doe"
    email: "jon_doe@example.com"
    password: "hardtoguess"
  })
  {
    email
    name
  }
}
```

### Create nooble mutation
```graphql
mutation($file: Upload!) {
createNooble(input: {
  title: "test_title"
  description: "this is a sample description"
  category: "people_and_blogs"
  creator: "jon_doe@example.com"
  file: $file
})
}
```
