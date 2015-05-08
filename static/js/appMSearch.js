var returnPage = null;
var count = 0 ;
$(function() {
    //获取当前页面的url,判断当前页面是搜索前还是搜索之后,搜索之后URL参数会变化
    //由于要搜索之后需要返回到搜索首页
    var currUrl = window.location.href;  
    if (currUrl.indexOf("c=")!=-1) {
        var len = currUrl.length;
        var cIndex = currUrl.indexOf("c=");
        returnPage = currUrl.substring(cIndex,len);
        //将一次跳入的页面写入cookie或者选择某种应用之后URL改变之后的值写入cookie
        document.cookie="c="+returnPage;
    };
    if(currUrl.indexOf("s=")!=-1){
        $("#inputText img").css({"display":"block"});
    }
    var search = location.search,searchVal,searchReg=/s=([^&]*)/g;
    $("#sure").click(function() {
        var isLogin = location.search,
            reg = /(?!isLogin=)true/g,
            name = $.trim($("#inputText input").val()),
            Surl;

        if (isLogin.match(reg)) {
            isLogin = 'true';
        } else {
            isLogin = 'false';
        }
        Surl = "/appMCenter/appMSearch.php?s=" + name + "&isLogin=" + isLogin;
        location.href = Surl;
    });
    //搜索框
    $(".search-ipt").on("click", function() {
        $(this).val('');
    });
    $(".search-btn").click(function() {
        s = $.trim($(".search-ipt").val());

        $(".search_input").css({
            'display': 'block'
        });
        if (s != "") {
            window.open("appMSearch.php?s=" + s);
        }
    });
    //搜索框回车功能
    $(".search_input input").focus(function() {
        $(this).keydown(function(e) {
            if (e.keyCode == 13) {
                if (document.all) {
                    event.returnValue = false;
                } else {
                    e.preventDefault();
                };

                $("#sure").click();
            }
        });
    });
    $(".appRecommend-l").mouseenter(function() {
        $(this).find(".floatDiv").fadeIn();
    });
    $(".appRecommend-l").mouseleave(function() {
        $(this).find(".floatDiv").fadeOut();
    });
    $(".appRecommend-r-inner").mouseenter(function() {
        $(this).children(".floatDiv").fadeIn();
    });
    $(".appRecommend-r-inner").mouseleave(function() {
        $(this).children(".floatDiv").fadeOut();
    });

    //菜单弹出框
    $(".menu").click(function(e) {
        $obj = $("#menuDrop");
        e.stopPropagation();
        
        $obj.hasClass('hide')?$obj.removeClass("hide"):$obj.addClass("hide");
    });

    $(document).click(function() {
        if (!$("#menuDrop").hasClass("hide")) {
            $("#menuDrop").addClass("hide");
        }
    });

    $("#menuDrop div").click(function(e) {
        var isLogin = location.search,
            reg = /(?!isLogin=)true/g,
            url;

        if (document.all) {
            e.returnValue = false;
        } else {
            e.preventDefault();
        }
        if (isLogin.match(reg)) {
            location.href = $(this).find("a").attr("href") + "&isLogin=true";
        } else {
            location.href = $(this).find("a").attr("href") + "&isLogin=false";
        }

    })

    // 搜索栏
    $(".seach_img").click(function() {
        $("#imgChange").removeClass("flower");
        $("#imgChange").addClass("return");
        $(".search_input").show();
        $(".search_input input").focus();
        $(".search").hide();
        $(".search").css({
            width: "6%",
            minWidth: "48px"
        });
        $(".useCenter").css({
            display: "none"
        });
    });

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

    $(".appMCenter_pop,.appMCenter_popCont").click(function(){
        $(".appMCenter_pop").hide();
        $(".appMCenter_popCont").hide();
    });

    $(".listC").click(function() {
        location.href = $(this).find(".listC_L a").attr("href");
    })
    
    //页面跳转
    var titleCenterchild = $.trim($("#titleCenter").text());

    if ( titleCenterchild == "云盘应用" || titleCenterchild == "雲盤應用") {
        $("#imgChange").css({
            'backgroundImage': 'url("/smarty/images/appMCenter/big_arrow_down.png ")',
            'backgroundSize': '26px 32px',
            'backgroundPosition': '5px 14px'
        })
    }

    $("#imgChange").click(function() {
        var isLogin = location.search,
            reg = /(?!isLogin=)true/g,
            name = $.trim($("#inputText input").val()),
            Surl;

        if (isLogin.match(reg)) {
            isLogin = 'true';
        } else {
            isLogin = 'false';
        }
        var cookie = document.cookie;   //取得cookie的中的值,如果没有默认是
        //回到android应用页面首页
        if(cookie!=null&&cookie!=undefined){
            var start = 0;
            var end  = 0 ;
            if(cookie.indexOf("c=")!=-1){
                start = cookie.indexOf("c=");
                end = cookie.indexOf(";",start+2);
                if(end==-1) {
                end=cookie.length; 
               }
                var currArg = cookie.substring(start+2,end); 
                Surl = "/appMCenter/appMSearch.php?"+currArg;

            }else{
               Surl = "/appMCenter/appMSearch.php?c=0&isLogin="+isLogin;
            }
        }else{
            Surl = "/appMCenter/appMSearch.php?c=0&isLogin="+isLogin;
        }
        location.href = Surl;
    });

    $('#search_close').click(function(){
        $(".search_input input").val('');
        $("#inputText img").css({"display":"none"});
    });

    if(search.indexOf('s=') != -1){
        var res = searchReg.exec(search);
        if(res&&res[1]!=''){
            $('#inputText input').val(decodeURI(res[1]));
            $('.seach_img').click();
            $("#titleCenter").text("搜索:'"+res[1]+"'");
        }else{
            $("#titleCenter").text("搜索:'全部應用    '");
        }
    }
     //更改搜索框的效果,当输入框没有值时后面的X图片不出现
     $("#searchKeyWord").on("keydown",function(){
           var inputVal = $.trim(document.getElementById("searchKeyWord").value);
           if (inputVal!=null&&inputVal!=undefined&&inputVal!="") {
            $("#inputText img").css({"display":"block"});
           }else{
             $("#inputText img").css({"display":"none"});
           }
     })

});