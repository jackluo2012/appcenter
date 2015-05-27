<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<!-- saved from url=(0039)http://localhost/appDetail.php?appid=41 -->
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	<title> {{i18n $.Lang "detail"}} </title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<meta content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0;" name="viewport" />
	<meta name="apple-touch-fullscreen" content="yes">
	<meta content="yes" name="apple-mobile-web-app-capable" />
	<meta content="telephone=no" name="format-detection" />
	<meta content="email=no" name="format-detection" />
	<script type="text/javascript" src="/static/js/jquery-1.7.2.min.js"></script>
	<script type="text/javascript" src="/static/js/appMDetail.js"></script>
	<link href="/static/css/owl.carousel.css" rel="stylesheet">
	<link href="/static/css/owl.theme.css" rel="stylesheet">
	<link rel="stylesheet" type="text/css" href="/static/css/Mstyle.css">
	<style type="text/css">
	</style>
</head>
<body style="background-color:#e3e8ee">
	<div class="main-oper">
		<div class="main_content clearfix">
				<div class="logo logo1"></div>
			<span class="fontColor">{{.AppInfo.Name}}</span>
		</div>
	</div>
	<div id="contener">
		<link rel="stylesheet" type="text/css" href="/static/css/appMDetail.css">
		<div class="main centerize">
			<div class="detail-main">
				<div class="app-part1 clearfix">
					<div class="app-l">
						<div class="app-img">
							<img src="{{.AppInfo.IconUrl}}"></div>
					</div>
					<div class="app-m clearfix">
						<div class="app-m-line1">
							<span class="app-name"> <b>{{.AppInfo.Name}}</b>
							</span>
						</div>
						<div class="app-m-line2">
						    <div id="appInfo1" style="margin-bottom:5px;">
						    <label class="app-downCout">{{i18n $.Lang "download"}} :&nbsp;<span class="download_count">{{.AppInfo.DownloadCounts}}</span>{{i18n $.Lang "downs"}} </label>
							<span class="app-size" style="margin-left:20px">{{i18n $.Lang "size"}} :&nbsp;{{.AppInfo.Size}}</span>
							<span class="app-version1" style="margin-left:20px">{{i18n $.Lang "version"}}:&nbsp;{{.AppInfo.Version}}</span>
							<span class="app-developer" style="margin-left:20px">{{i18n $.Lang "developer"}}:&nbsp;{{.AppInfo.Author}}</span><br/>
						    </div>
							<span class="app-update">{{i18n $.Lang "updatetime"}}:&nbsp;{{.AppInfo.Created}}</span>
						</div>
						<div class="app-m-line3">
							{{if eq .AppInfo.Install "3"}}
							
							<a class="downloadUrl download_Button" target="_blank" href="{{.AppInfo.DownLoadUrl}}" >{{i18n $.Lang "update"}} </a>
							{{else if eq .AppInfo.Install "2"}}
							<a class="downloadUrl download_Button" href="#" target="_blank">{{i18n .Lang "readydown"}}</a>
							{{else}}
								<a class="downloadUrl download_Button" target="_blank" href="{{.AppInfo.DownLoadUrl}}" >{{i18n $.Lang "download"}}</a>
							{{end}}

						</div>
					</div>
				</div>
				<div class="app-part2">
					<div class="part2-title"></div>
					<span class="fontSize">{{i18n $.Lang "description"}}</span>
					<div class="margi">
						<div class="part2-description">{{.AppInfo.Description}}</div>
						<div class="app-carousel">
			
			<!-- loop---kitty 2014--09-28- end-->
		

			<div class="row">
				<div class="span12">

					<div id="owl-demo" class="owl-carousel">
						{{range $k, $v := .Screens}}
							{{if $v}}
								<div class="item">
									<img src="{{$v}}" alt="Owl Image">
								</div>
							{{end}}
						{{end}}
						
					</div>

				</div>
			</div>

			<!-- kitty--2014-09-28-14:25 start -->


		</div>
	</div>
</div>

</div>
</div>
</div>
 <!-- 遮罩-kitty-2014-10-20 -->
  <div class="appMCenter_pop" style="display:none;">
   </div>
   <section  class="appMCenter_popCont" style="display:none;">
      <div class="appMCenter_popConttitile">{{i18n $.Lang "loginaccess"}}</div>
   </section>
<script type="text/javascript" src="/static/js/owl.carousel.js"></script>
<script src="/static/js/application.js"></script>
<script>
    $(document).ready(function() {
      $("#owl-demo").owlCarousel({
        autoPlay: 3000,
        items :2
        // itemsDesktop : [1199,3],
        // itemsDesktopSmall : [979,3]
      });

    });
</script>
<script type="text/javascript">
//*

//弹出层
$(function(){

  $(".appMCenter_popContBt button").click(function() {                
                     $(".appMCenter_pop").hide();
                     $(".appMCenter_popCont").hide();
  });
 //下载前登录验证
    $(".download_Button").click(function(e){
        //*
        var isLogin = location.search,
            reg = /secretkey=(*){32}/g;

        if(!e.preventDefault()){
            e.returnValue = false;
        }
        //*
        if(isLogin.match(reg)){
            location.href = $(this).attr("href");
        }else{
               $(".appMCenter_pop").show();
                $(".appMCenter_popCont").show();

        }
    });


})

//*/

    </script>
</body>
</html>