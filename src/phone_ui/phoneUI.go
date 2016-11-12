package phone_ui

import (
	"html/template"
)

var IndexTpl= template.Must(template.New("PhoneUI_Index").Parse(`<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>支付通讯录</title>
    <meta name="viewport" content="initial-scale=1, maximum-scale=1">
    <link rel="shortcut icon" href="/favicon.ico">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">

    <link rel="stylesheet" href="http://g.alicdn.com/msui/sm/0.6.2/css/sm.min.css">
    <link rel="stylesheet" href="http://g.alicdn.com/msui/sm/0.6.2/css/sm-extend.min.css">
 </head>
<body>


<div class="page-group">

   <div class="page page-current" id="mainPage">

	<header class="bar bar-nav">
	  <span class="icon icon-friends">&nbsp通讯录</span>
	  
	</header>
	<div class="bar bar-header-secondary">
	  <div class="searchbar">
 		
	    <a class="searchbar-cancel">取消</a>
	    <div class="search-input">
	      <label class="icon icon-search" for="search"></label>
	      <input type="search" id='search' placeholder='输入关键字...'/>
	    </div>
	  </div>
	</div>
	
 
	<div class="content" id="allcontact">
	  <div class="list-block contacts-block">

{{range .DeptList}}

	    <div class="list-group">
	      <ul>
		<li class="list-group-title">{{.DeptName}}</li>
		{{range .PList}}
		<li>
		 <a href="#{{.PId}}">
		  <div class="item-content">
		    <div class="item-inner">
		      <div class="item-title">{{.Name}}</div>
		    </div>
		  </div>
		 </a>
		</li>

		{{end }}

	      </ul>
	    </div>
{{end}}	  
	   
	  </div>
	</div>

　　　　<div class="content" id="searchResult">
      <div class="list-block contacts-block">
        <div class="list-group">
          <ul>
		<li class="list-group-title"　id="searchResultListGroup">搜索结果</li>
		<li>
		 <a href="#router2" ><!--搜索结果也直接链接到具体的详情页，不需要另行请求-->
		  <div class="item-content">
		    <div class="item-inner">
		      <div class="item-title">请输入信息搜索</div>
		    </div>
		  </div>
		 </a>
		</li>
            </ul>
	   </div>
          </div>		
    </div> 


    </div>


<!-- for detail info -->  
{{range .allPList}}

	<div class="page" id='{{.PId}}'>
	  <header class="bar bar-nav">
	    <a class="button button-link button-nav pull-left back" href="">
	      <span class="icon icon-left"></span>
	      返回
	    </a>
	    <h1 class='title'>{{.Name}}</h1>
	  </header>
	  <div class="content">
	    <div class="list-block">
		<ul>
			<li>
				<div class='item-content'>
					<div class='item-media'><i class="icon icon-card"></i></div>
					<div class='item-inner'>
						<div class="item-title label">姓名</div>
						<div class="item-input">
							<input type="text" > {{.Name}} </input>
						</div>
						
					</div>
				</div>
			</li>
			<li>
				<div class='item-content'>
					<div class='item-media'><i class="icon icon-app"></i></div>
					<div class='item-inner'>
						<div class="item-title label">部门</div>
						<div class="item-input">
							<input type="text" >{{.DeptName}}</input>
						</div>
						
					</div>
				</div>
			</li>
			<li>
				<div class='item-content'>
					<div class='item-media'><i class="icon icon-message"></i></div>
					<div class='item-inner'>
						<div class="item-title label">邮箱</div>
						<div class="item-input">
							<a href="mailto:{{.Mail}}"><input type="text" >{{.Mail}}</input></a>
						</div>
						
					</div>
				</div>
			</li>
			<li>
				<div class='item-content'>
					<div class='item-media'><i class="icon icon-phone"></i></div>
					<div class='item-inner'>
						<div class="item-title label">电话</div>
						<div class="item-input">
							<input type="text" >{{.Phone}}</input>
						</div>
					</div>
				</div>
			</li>
			<li>
				<div class='item-content'>
					<div class='item-media'><i class="icon icon-phone"></i></div>
					<div class='item-inner'>
						<div class="item-title label">固话</div>
						<div class="item-input">
							<input type="text" >{{.Tel}}</input>
						</div>
					</div>
				</div>
			</li>
		</ul>
	    </div>
	   <div class="content-block">
		<div class="row">
			<div class="col-50"><a href="#" class=" back button button-big button-fill button-danger">取消</a></div>
                        <div class="col-50"><a href="tel:{{.Phone}}" class="button button-big button-fill button-success">拨打</a></div>
		</div>
	   </div>
	  </div>
	</div>

{{end}}

<!-- 预留的关于
<div class="popup popup-about">
  <div class="content-block">
    <p>About</p>
    <p><a href="#" class="close-popup">Close popup</a></p>
    <p>Create by Shanggl</p>
  </div>
</div>


</div>
-->


   <script type='text/javascript' src='http://g.alicdn.com/sj/lib/zepto/zepto.min.js' charset='utf-8'></script>
    <script type='text/javascript' src='http://g.alicdn.com/msui/sm/0.6.2/js/sm.min.js' charset='utf-8'></script>
    <script type='text/javascript' src='http://g.alicdn.com/msui/sm/0.6.2/js/sm-extend.min.js' charset='utf-8'></script>
   <script>
	//init 
	
	$(document).on('ready',function(e){
		$('#searchResult').hide()
	
	});
	
	$('#search').on('focus',function(e){
		$('#allcontact').hide()
		$('#searchResult').show()
	
	});

	$('#search').on('blur',function(e){
		//需要先判断是否有搜索结果，没有的话就自动返回
		var v=$('#search').val()
		if(""==v){
			$('#allcontact').show()
			$('#searchResult').hide()
		}
	
	});
	

	$('#search').on('change',function(e){
		var v=$('#search').val()
		
		if (""==v ){
			$("#allcontact").show()
			$("#searchResult").hide()
		}
		//ajax 请求搜索内容
		//alert("change"+v);
		
	
	});


	$(".searchbar-cancel").on('click',function(e){
	        $('#search').val("");
		$("#allcontact").show()
		$("#searchResult").hide()
		
	});
	

	$.init();
  </script>

  </body>
</html>




`))

