version: "0.1"
database:
  dsn : "host=aws-0-us-west-1.pooler.supabase.com user=postgres.xqpbqcikgjzzaleftrac password=Seek15wish16@ port=5432 dbname=postgres"
  db  : "postgres"
  outPath :  "./src/Database/Dao"
  outFile :  "dao.go"
  tables : ["Tasks"]
  fieldNullable : false
  fieldWithIndexTag : false
  fieldWithTypeTag  : false