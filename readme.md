#  file encryption and send to Pastebin

this is the simple program that testing for detection to taking out file 

change this section to your pastebin API key
```
// Pastebin API setting
const (
	pastebinAPIURL   = "https://pastebin.com/api/api_post.php"
	pastebinAPIKey   = "<YOUR API KEY>" // Pastebin API Key
	pastebinUserKey  = ""                     // option <login users key>
	pastebinPrivate  = "1"                    // 0=pub, 1=invisible, 2=private
	pastebinOption   = "paste"                // API option; paste
)


```


build:
  ```
  go build encode.go
  go build decode.go
  ```

usage:
  ```
[encode and send]
  ./encode <file>

[decode and get]
  ./decode https://pastebin.com/<PATH> <output file>
  ```

example:

![image](https://github.com/user-attachments/assets/92745951-62ca-4dda-afc9-ecbc0a8ad58f)

![image](https://github.com/user-attachments/assets/5833914e-24dc-430d-a1f1-74a7718a9402)

![image](https://github.com/user-attachments/assets/e6ebd9df-e52d-4fe5-8be8-69f93952a73a)




