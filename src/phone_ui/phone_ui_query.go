package phone_ui

import (
	"html/template"
)

var QueryResultTpl= template.Must(template.New("PhoneUI_Index_Query").Parse(`
<div class="list-block contacts-block">
 <div class="list-group">
    <ul>
	<li class="list-group-title">查询结果</li>
	{{range .pList}}
	<li>
	 <a href="/p?pid={{.PId}}">
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
`))

