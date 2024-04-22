# HASH_SCAN

Hash Scan liest den MD5 Hash saemtlicher Dateie sich rekursive befindlich unter dem initialen Verzeichnis aus.
Hash Check checkt die MD5 Hashes eines Logs gegen die Hash-DB von Abuse.ch.


```
go run main.go -o output.txt C:\Dell\
```

Beispiel:
```
>go run main.go -o output.txt C:\Drivers
File: C:\Drivers\474.14-quadro-rtx-desktop-notebook-win10-win11-64bit-international-dch-whql.exe, MD5 Hash: 5d892e2d6dba35a3dfc131936747bd5d
File: C:\Drivers\DellCommandUpdateApp_Setup.exe, MD5 Hash: 41b2c34ea924541eaf5f5df09bbea969
File: C:\Drivers\Mup.xml, MD5 Hash: f356f25db871cb8664bf1d4f019150b7
File: C:\Drivers\package.xml, MD5 Hash: 96ae0f840a6ab1e5b17fe28777aae3f0
File: C:\Drivers\474.14-quadro-rtx-desktop-notebook-win10-win11-64bit-international-dch-whql.exe, MD5 Hash: 5d892e2d6dba35a3dfc131936747bd5d
File: C:\Drivers\DellCommandUpdateApp_Setup.exe, MD5 Hash: 41b2c34ea924541eaf5f5df09bbea969
File: C:\Drivers\Mup.xml, MD5 Hash: f356f25db871cb8664bf1d4f019150b7
File: C:\Drivers\package.xml, MD5 Hash: 96ae0f840a6ab1e5b17fe28777aae3f0
```



# CHECK HASHES
Donwload https://bazaar.abuse.ch/export/csv/full/

Entpacke ZIP das full.csv im Directory liegt

```
go run main.go -o output.txt C:\Dell
```
```
go run check_hash.go -i output.txt
```
