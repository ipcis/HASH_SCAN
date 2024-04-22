# HASH_SCAN
```
go run main.go -o output.txt C:\Dell\
```



# CHECK HASHES
Donwload https://bazaar.abuse.ch/export/csv/full/

Entpacke ZIP das full.csv im Directory liegt

go run main.go -o output.txt C:\Dell

go run check_hash.go -i output.txt
