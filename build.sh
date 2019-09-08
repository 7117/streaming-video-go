# webserver 进行编译出来  再移动到真正webui下面  将template全部移动过去
# 我们需要的是  web的编译文件+页面资源文件  放在bin文件夹下即可
# webserver(web的编译体+template)

cd ~/work/src/github.com/avenssi/video_server/web
go install
cp ~/work/bin ~/work/bin/video_server_web_ui/web
cp -R ~/work/src/github.com/avenssi/video_server/template  ~/work/bin/video_server_web_ui

