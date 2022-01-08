### MyTEnS & My Carrier Test

#### Stacks
 - Go (go version go1.17.1 linux/amd64)
 - Commando (github.com/thatisuday/commando)

#### How to run this app
````
$ git clone https://github.com/emixbal/MyTEnsApp.git
$ cd MyTEnsApp
$ go build
````

#### Use the  app

 - show help
	 ````
	 $ ./MyTEnsApp -h
	 ````
	 ![show help](https://raw.githubusercontent.com/emixbal/MyTEnsApp/main/images/help.png)
 - copy file log 
	 ````
	 $ ./MyTEnsApp <file name> -t <opsional, valid: text, json> -o <opsional, valid: valid path>
	 ````
	##### Contoh

	 ![example](https://raw.githubusercontent.com/emixbal/MyTEnsApp/main/images/tanpa%20t%20tanpa%20o.png)

	 ![example](https://raw.githubusercontent.com/emixbal/MyTEnsApp/main/images/dengan%20o%20tanpa%20t.png)

	 ![example](https://raw.githubusercontent.com/emixbal/MyTEnsApp/main/images/tanpa%20o.png)

	 ![example](https://raw.githubusercontent.com/emixbal/MyTEnsApp/main/images/tanpa%20t%20tanpa%20o.png)

	 ![example](https://raw.githubusercontent.com/emixbal/MyTEnsApp/main/images/tanpa%20t%20tanpa%20o.png)
	 
	##### Contoh Validasi

	 ![example](https://raw.githubusercontent.com/emixbal/MyTEnsApp/main/images/validasi%20file%20ext%20tujuan%20berbeda%20dangan%20t.png)

	 ![example](https://raw.githubusercontent.com/emixbal/MyTEnsApp/main/images/validasi%20hanya%20bisa%20accept%20json%20dan%20text.png)

	 ![example](https://raw.githubusercontent.com/emixbal/MyTEnsApp/main/images/validasi%20invalid%20patht.png)
