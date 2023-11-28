# bing-wallpaper
### Download wallpaper from bing.com
#### Usage by default 
```shell
$ ./bing-wallpaper get
2023/11/28 18:27:25 Save wallpaper image to ./bing-wallpapers
2023/11/28 18:27:26 Response status: 200 OK
2023/11/28 18:27:26 http://bing.com/th?id=OHR.AssiniboineProvincialPark_ROW9470886401_UHD.jpg
2023/11/28 18:27:26 Response status: 200 OK
2023/11/28 18:27:28 INFO: Save wallpaper ./bing-wallpapers/OHR.AssiniboineProvincialPark_ROW9470886401_UHD.jpg
```

#### Usage with command line arguments
```shell
$ ./bin/bing-wallpaper get --res 1920x1080 --target any-bing-wallpapers  
2023/11/28 18:31:27 Save wallpaper image to any-bing-wallpapers
2023/11/28 18:31:28 Response status: 200 OK
2023/11/28 18:31:28 http://bing.com/th?id=OHR.AssiniboineProvincialPark_ROW9470886401_1920x1080.jpg
2023/11/28 18:31:28 Response status: 200 OK
2023/11/28 18:31:28 INFO: Save wallpaper any-bing-wallpapers/OHR.AssiniboineProvincialPark_ROW9470886401_1920x1080.jpg
```

