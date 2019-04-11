[json-generator](https://www.json-generator.com/)

```go
[
  '{{repeat(5, 7)}}',
  {
    _id: '{{objectId()}}',
    id: '{{index(1)}}',
    firstname: '{{firstName()}}',
    lastname: '{{surname()}}',
    info: {
        city: '{{city()}}',
        phone: '{{integer([0], [99999])}}'
    }
  }
]
```
