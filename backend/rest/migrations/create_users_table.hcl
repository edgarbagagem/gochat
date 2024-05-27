table "users" {
  schema = schema.main
  column "id" {
    null = true
    type = integer
  }
  column "username" {
    null = false
    type = text
  }
  column "password" {
    null = false
    type = text
  }
  column "photo" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
}
schema "main" {
}