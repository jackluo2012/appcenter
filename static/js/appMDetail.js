/**
 * Created on 14-7-10.
 */
;
(function(window, document) {
 
    var UrlParas = function(url) {
        return UrlParas.fn.init(url);
    }
    UrlParas.VERSION = '1.0.0';
    UrlParas.fn = UrlParas.prototype = {
 
        url: "",
        pathname: "",
        paras: "",
        init: function(url) {
            this.url = url;
            this.pathname = url.split("?")[0];
            this.paras = this.get();
            return this;
        },
 
        //以object类型返回url参数及其取值
        get: function(option) {
            var paraStr, paras,
                url = this.url;
            if (url) {
                paraStr = url.split("?")[1];
                if (paraStr) {
                    paras = {};
                    paraStr = paraStr.split("&");
                    for (var n in paraStr) {
                        var name = paraStr[n].split("=")[0];
                        var value = paraStr[n].split("=")[1];
                        paras[name] = value;
                    }
                } else {
                    return {};
                }
                if (!option) {
                    return paras;
                } else {
                    return paras[option] ? paras[option] : "";
                }
 
 
            }
        },
 
        //重设url参数取值，若无此参数则进行创建,若参数赋值为null则进行删除
        set: function(option) {
            var i, name, val;
            if (arguments.length == 2) {
                name = arguments[0];
                val = arguments[1];
                option = {
                    name: val
                };
            }
            if ("string" === typeof option) {
                this.paras[option] = "";
            } else if ("object" === typeof option) {
                for (i in option) {
                    if (option[i] === null) {
                        delete this.paras[i];
                    } else {
                        this.paras[i] = option[i];
                    }
                }
            } else {
 
            }
            return this.build();
        },
 
        //删除url中指定参数返回新url
        remove: function(option) {
            var i;
            if ("string" === typeof option) {
                option = option.split(",");
                for (i in option) {
                    delete this.paras[option[i]]
                }
 
            }
            return this.build();
        },
 
        //根据url和处理过的paras重新构件url
        build: function() {
            var i,
                newUrl = this.pathname + "?";
 
            for (i in this.paras) {
                newUrl += (i + "=" + this.paras[i] + "&");
            }
 
            return newUrl.substr(0, newUrl.length - 1);
        }
 
 
    }
 
    UrlParas.fn.init.prototype = UrlParas.fn;
 
    window.urlParas = UrlParas;
 
})(window, document);


$(function(){
    //如果是手机或平板等,应用的大小,下载量等信息换行显示
    var clientWidth = document.body.clientWidth;
    if(clientWidth<=599){
        $(".app-downCout").after("<br/>");
        $(".app-size").after("<br/>");
        $(".app-version1").after("<br/>");
        $(".app-downCout").css({"margin-left":"0px","margin-bottom":"2px"});
        $(".app-size").css({"margin-left":"0px","margin-bottom":"2px"});
        $(".app-version1").css({"margin-left":"0px","margin-bottom":"2px"});
        $(".app-developer").css({"margin-left":"0px","margin-bottom":"2px"});
        $("#appInfo1").css({"margin-bottom":"2px"})
        $(".main_content .logo").css({"margin-left":"20px"})
    }
    if(clientWidth<1000&&clientWidth>=600){
        $(".app-size").css({"margin-left":"10px","margin-bottom":"2px"});
        $(".app-version1").css({"margin-left":"10px","margin-bottom":"2px"});
        $(".app-developer").css({"margin-left":"10px","margin-bottom":"2px"});
        $("#appInfo1").css({"margin-bottom":"2px"})
        $(".main_content .logo").css({"margin-left":"70px"})
    }
    if(clientWidth>1024){
        $(".main_content .logo").css({"margin-left":"100px"})
    }
	if(urlParas(location.href).get("isLogin") !== "true"){

		$(".main_content a").attr("href",urlParas("appMSearch.php").set({"isLogin":"false"}));
	}else{
		$(".main_content a").attr("href",urlParas("appMSearch.php").set({"isLogin":"true"}));
	}


	var success,thisname,mark=0;
	
    $(".seach_img").click(function(){
        mark++;
        if(mark%2==0){
            $(".search_input").css({'display':'block'});
        }else{
            $(".search_input").css({'display':'none'});
        }
    });
	
	thisname = $(".app-name b").text();
	$(".main_content span").text(thisname);


	function isLogin(){
		$.ajax({
			type:'GET',
			url:"common/isLogin.php",
			data:"",
			async: false,
			success:function(result){
			
			if(result == "logout"){
				success = 0;
			}else{
				success = 1;
			}
		}
		});
		return success;
	}
	var p = 1;
	linum = $(".app-comments >ul >li").length;
		if(linum<10){ 
		$(".button-more").unbind( "click" )
		$(".button-more").html("抱歉..没有更多的了。");
	 }

	 $();

    //搜索框
	$(".search-ipt").one("click",function(){
		$(this).val('');
	});
	
	$(".search-btn").click(function(){
		s = $.trim($(".search-ipt").val());
	    if(s != ""){
	    	window.open("appSearch.php?s="+s);
	    }
	});
    //搜索框回车功能

    $(".search-ipt").focus(function(){
        $(this).keydown(function(e){
        if(e.keyCode==13){
			if (document.all) {
         		event.returnValue = false;
        	} else {
         		e.preventDefault();
        	};
		    $(".search-btn").click();
		}   		
        });
       	
    });


    //分类按钮	
	$("html").not($(".dropdown-menu")).click(function(){
		if($(".dropdown-menu").css("display")=="block"){
			$(".more-btn").click();
		}
	});

	$(".more-btn").toggle(function(){
		$(".dropdown-menu").fadeIn();	
	},function(){
		$(".dropdown-menu").fadeOut();
	});	

	object = $(".catePane-phb ul li");
	object.not(object.first()).addClass("simple").children().not($(".app-name,.app-rank")).hide();

	object.not(object.first()).mouseenter(function(){

				$(this).removeClass("simple").addClass('curr').children().show();
				object.first().removeClass("curr").addClass("simple").children().not($(".app-name,.app-rank")).hide();
	});
	object.mouseleave(function(){
				$(this).removeClass("curr").addClass('simple').children().show();
				object.first().removeClass("simple").addClass("curr").children().show();
				object.not(object.first()).addClass("simple").children().not($(".app-name,.app-rank")).hide();
	});	

	//ajax返回更多评论
	$(".button-more").click(function(e){

		if (document.all) {
         event.returnValue = false;
        } else {
         e.preventDefault();
        };
		var appid = $(".button-more").attr("data-appid"); 
		data = {"appid":appid,"p":p};
		handleUrl = "moreComments.php";
		$.get(handleUrl,data,function(result){
			num = $(result).find("li").length;
			if (num <= 0){

			}else{
				$(".app-comments").append(result);
				p++;
				if(num<=10){ 
					$(".button-more").unbind( "click" )
					$(".button-more").html("抱歉..没有更多的了。");

				 }
			}
		},'html');	
	});

//发表评论ajax
$(".comment-submit span").click(function(){
	handleUrl = "appMCommentHandle.php";
	text  = $.trim($(".comment-submit textarea").val());
	appid = $(".comment-submit").attr("data-appid");
	account = $(".comment-submit").attr("data-account");
	data = {"text":text,"appid":appid,"account":account};
	console.log(data);
	$.post(handleUrl,data,function(result){
		if(result.status == 1){console.log("发表成功");window.location="appMDetail.php?appid="+appid;}
	},"json");
});

//回复评论ajax
$(".app-comments").on("click",'.replay1,.replay2',function(e){
	if(isLogin() == 0){alert("<!--###抱歉亲，请先登录再回复评论吧###-->") ;return false;}
	if (document.all) {
         event.returnValue = false;
        } else {
         e.preventDefault();
        };
    html = 	"<li>"
    		+"<div class='replay-zone'><textarea class = 'replay-text'></textarea><span class='replay-submit button-azury'"
    		+"data-reuserId='"+$(this).attr("data-userId")+"' data-reuserName='"+$(this).attr("data-userName")+"' data-originalId='"+$(this).attr("data-originalId")+"' "
    		+">发表</span></div>"
    		+"</li>";
    if($(this).hasClass("replay1")){
    	object = $(this).parents(".comment-zone").siblings('ul');    		
    }else{
    	object = $(this).parents(".recomment-zone").parent().parent('ul');
    }
   	if(object.has("li .replay-zone").length == 0){
		object.append(html);		
		object.find(".replay-zone").hide().slideDown();
   	}else{
   		object.find(".replay-zone").slideUp(function(){
   			$(this).remove();
   			object.append(html);
			object.find(".replay-zone").hide().slideDown();	
   		});		
   	}
});

    $(".logo1").click(function(){
        console.log("hello");
        window.history.go(-1);
    });

$(".app-comments").on('click','.replay-submit',function(){
	$this = $(this)
	appid = $(".comment-submit").attr("data-appid"); 
	data ={
		"appid":appid,
		"original_comment_id":$(this).attr("data-originalId"),
		"userid":"19",
		"username":"Mic",
		"reuserid":$(this).attr("data-reuserId"),
		"reusername":$(this).attr("data-reuserName"),
		"comment":$(this).siblings('textarea.replay-text').val()
	} 
	handleUrl = "appReplay.php";
	$.post(handleUrl,data,function(result){
		if(result.status == 1){
			// console.log("回复成功");window.location="appDetail.php?appid="+appid;
			html = 	"<div class='recomment-zone'>"
					+"<div class='user-icon'><img src='"+result.replay[0].face_url+"' alt=''></div>"
					+"<div class='user-recomment'>"
					+"<div class='comment-line1'><a href='myspace.php?account="+result.replay[0].account1+"'>"+result.replay[0].username+"</a> 回复 <a href='myspace.php?account="+result.replay[0].account2+"'>"+result.replay[0].reusername+"</a> 说："+result.replay[0].comment+"</div>"
					+"<div class='comment-line2'>"
					+"<div class='comment-date'>"+result.replay[0].created+"</div>"
					+"<a class='replay2' data-userId='"+result.replay[0].userid+"' data-userName='"+result.replay[0].username+"' data-originalId='"+result.replay[0].original_comment_id+"' href=''>回复</a>"
					+"</div></div></div>";
			$this.parent().parent().html(html).hide().fadeIn();
		}
	},"json");




});
//图片区轮播效果
    var parLn=$(".screen-zone").width();//获取他父亲元素的宽度
    var $li1 = $(".screenImg .screen-array");
    var lilen=$li1.length;
    var $window1 = $(".app-carousel .screenImg");
    var	$left1 = $(".switch-left");
    var	$right1 = $(".switch-right");
   //定义图片的宽度
    $li1.css("width",100/lilen+"%");//定义图片的宽度
    var animateLeft=null;
    switch (true){
        case parLn>480:
            animateLeft=100/2+"%";
            $window1.css("width", 100*lilen/2+"%");//一排显示三张图片
            break;
        case parLn<=480&&parLn>240:
            animateLeft=100/2+"%";
            $window1.css("width", 100*lilen/2+"%");//一排显示2张图片
            break;
        case parLn<=240:
            animateLeft=100+"%";
            $window1.css("width", 100*lilen+"%");//一排显示1张图片
            break;
    }

	var lc1 = 0;
	var rc1 = $li1.length-3;
	
	$left1.click(function(){
		if (lc1 < 1) {
			alert("<!--###已经是第一张图片###-->");
			return;
		}
		lc1--;
		rc1++;
		$window1.animate({left:'+='+animateLeft}, 1000);
	});

	$right1.click(function(){
		if (rc1 < 1){
			alert("<!--###已经是最后一张图片###-->");
			return;
		}
		lc1++;
		rc1--;
		$window1.animate({left:'-='+animateLeft}, 1000);
	});	



var brTest=$(".part2-description").find("br");
$(brTest).before('</br>')


/*
    //下载前登录验证
    $(".download_Button").click(function(e) {
        var isLogin = location.search,
            reg = /(?!isLogin=)true/g;

        if (!e.preventDefault()) {
            e.returnValue = false;
            e.stopPropagation()
        }
        if (isLogin.match(reg)) {
            var appid = $(this).data("id"),
                value = parseInt($(".download_count[data-id=" + appid + "]").text());
            location.href = $(this).attr("href");
            $(".download_count[data-id=" + appid + "]").text(value + 1);
        } else {
            $(".appMCenter_pop").show();
            $(".appMCenter_popCont").show();
            setTimeout(function(){
                $(".appMCenter_pop").hide();
                $(".appMCenter_popCont").hide();
            },1000)
        }
    });
*/

});
//弹出层
/*
$(".appMCenter_popContBt button").click(function(){       

     $(".appMCenter_pop").hide();
     $(".appMCenter_popCont").hide();

})
*/