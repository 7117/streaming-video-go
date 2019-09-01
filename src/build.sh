# webserver 进行编译出来  再移动到真正webui下面  将template全部移动过去
# 我们需要的是  web的编译文件+页面资源文件  放在bin文件夹下即可
# webserver(web的编译体+template)


cd D:/Github/Streaming-video/src/web
go install
cp D:/Github/Streaming-video/bin D:/Github/Streaming-video/bin/ui/web
cp -R D:/Github/Streaming-video/src/template D:/Github/Streaming-video/bin/ui

