package phone_ui

import (
	"html/template"
)

var PersonTpl= template.Must(template.New("PhoneUI_Index_Person").Parse(`<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{.person.Name}}</title>
    <meta name="viewport" content="initial-scale=1, maximum-scale=1">
    <link rel="shortcut icon" href="/favicon.ico">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">

    <link rel="stylesheet" href="http://g.alicdn.com/msui/sm/0.6.2/css/sm.min.css">
    <link rel="stylesheet" href="http://g.alicdn.com/msui/sm/0.6.2/css/sm-extend.min.css">
 </head>
<body>
<!-- for detail info -->  


	<div class="page" id='{{.person.PId}}'>
	  <header class="bar bar-nav">
	    <a class="button button-link button-nav pull-left back" href="">
	      <span class="icon icon-left"></span>
	      返回
	    </a>
	    <h1 class='title'>{{.person.Name}}</h1>
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
							<input type="text" > {{.person.Name}} </input>
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
							<input type="text" >{{.person.DeptName}}/{{.person.TeamName}}/{{.person.Position}}</input>
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
							<a href="mailto:{{.person.Mail}}"><input type="text" >{{.person.Mail}}</input></a>
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
							<input type="text" >{{.person.Phone}}</input>
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
							<input type="text" >{{.person.Tel}}</input>
						</div>
					</div>
				</div>
			</li>
		</ul>
	    </div>
	   <div class="content-block">
		<div class="row">
			<div class="col-50"><a href="#" class=" back button button-big button-fill button-danger">取消</a></div>
                        <div class="col-50"><a href="tel:{{.person.Phone}}" class="button button-big button-fill button-success">拨打</a></div>
		</div>
	   </div>
	  </div>
	</div>

   <script type='text/javascript' src='http://g.alicdn.com/sj/lib/zepto/zepto.min.js' charset='utf-8'></script>
    <script type='text/javascript' src='http://g.alicdn.com/msui/sm/0.6.2/js/sm.min.js' charset='utf-8'></script>
    <script type='text/javascript' src='http://g.alicdn.com/msui/sm/0.6.2/js/sm-extend.min.js' charset='utf-8'></script>

</body>
</html>

`))

