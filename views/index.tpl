<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <title>{{i18n .Lang "appname"}}</title> 
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
  <meta content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0;" name="viewport" />
  <meta name="apple-touch-fullscreen" content="yes">
  <meta content="yes" name="apple-mobile-web-app-capable" />
  <meta content="telephone=no" name="format-detection" />
  <meta content="email=no" name="format-detection" />
  <link rel="stylesheet" href="/static/css/reset.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/appMSearch.css">
  <style type="text/css">
</style>
</head>
<body style="background-color:#e3e8ee">
  <div class="main-oper">
    <div class="operCont clearfix">
      <div class="operCont_left" >
        <div class="flower" id="imgChange"></div>
        <span class="useCenter" id="titleCenter">
        {{if eq .ctype "0"}}  
          {{i18n .Lang "androidapp"}}
        {{else}}
          {{i18n .Lang "webapp"}}  
        {{end}}

        </span>
      </div>
      
      <div class="search_input"  style ="display:none" id="search_input">
        <div id="inputText">
          <input  id="searchKeyWord" type="text" placeholder=""/>
          <img id='search_close' src="/static/img/search_clear.png" alt="" />          
        </div>
        <img src="/static/img/search_down.png" alt="" id="sure" style="width: 45px ;height: 45px;position: absolute;top: 8px;cursor: pointer;right:35px" />
      </div>

      <div class="search">
        <div class="menu"></div>
        <div class="seach_img" ></div>
      </div>
      <div class="menuDrop hide" id="menuDrop">
        <div>
          <span>
            <a href="?uid={{$.uid}}&udid={{$.udid}}&secretkey={{$.secretkey}}&c=0">
             {{i18n .Lang "androidapp"}} </a>
          </span>
        </div>
        <div>
          <span>
            <a href="?uid={{$.uid}}&udid={{$.udid}}&secretkey={{$.secretkey}}&c=1">
               {{i18n .Lang "webapp"}}  </a>
          </span>
        </div>
      </div>
  </div>
  </div>
<div class="contener" id="contenter">
  {{range .Apps}}

  <div class="listC clearfix">
    <div class="listC_L">
      <a href="/show/{{.Appid}}">
        <img src="{{.IconUrl}}" alt=""/>
      </a>
    </div>

    <div class="listC_M">
      <p class="size1">
        <span >
          <a href="/show/{{.Appid}}?uid={{$.uid}}&udid={{$.udid}}&secretkey={{$.secretkey}}">{{.Name}}</a>
        </span>
      </p>
      <p class="size3">
        <label>{{.Size}}</label>
        <label>
          {{i18n $.Lang "download"}}
          <b class="download_count" data-id="{{.Appid}}">{{.DownloadCounts}}</b>
          {{i18n $.Lang "downs"}} {{.Category}}</label>
      </p>
    </div>
  </div>
  {{if eq .Install "3"}}
      <a class="download_Button" href="{{.DownLoadUrl}}" target="_blank">
        <div class="listC_R">
          <span>
           {{i18n $.Lang "update"}} 
            </span>
        </div>
      </a>
      {{else if eq .Install "2"}}  
        <div class="listC_R">
          <span>
            {{i18n .Lang "readydown"}} 
            </span>
        </div>
      {{else }}    
       <a class="download_Button" href="{{.DownLoadUrl}}" target="_blank">
        <div class="listC_R">
          <span>
            {{i18n $.Lang "download"}} 
            </span>
        </div>
      </a>
      {{end}}
{{end}}

</div>
<script type="text/javascript" src="/static/js/jquery-1.7.2.min.js"></script>
<script type="text/javascript" src="/static/js/smoke.js"></script>
<script type="text/javascript" src="/static/js/appMSearch.js"></script>

</body>
</html>