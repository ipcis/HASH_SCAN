# HASH_SCAN

Hash Scan liest den MD5 Hash saemtlicher Dateie sich rekursive befindlich unter dem initialen Verzeichnis aus.
Hash Check checkt die MD5 Hashes eines Logs gegen die Hash-DB von Abuse.ch.


```
go run main.go -o output.txt C:\Dell\
```



# CHECK HASHES
Donwload https://bazaar.abuse.ch/export/csv/full/

Entpacke ZIP das full.csv im Directory liegt

go run main.go -o output.txt C:\Dell

go run check_hash.go -i output.txt
