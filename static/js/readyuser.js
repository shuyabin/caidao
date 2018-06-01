(function(){
    var $$$;
    var app = {
        togglePassVisibility : function(){
            var pass = $$$("pass_field");
            var value = pass.getValue();
            pass.config.type = this.getValue()?"text":"password";
            pass.refresh();
            pass.setValue(value);
        },
        saveUserForm : function(){
            var url = $$$("formpanel1").getValue();
            console.log(url);
            // $.ajax({
            //     url:"/addurl",
            //     type:'POST',
            //     success:function(msg){

            //     }
            // });
            
            $$$("userform").save();
        },
        addUser : function(){
            var list = $$$("userlist");
            var nid = list.add({
                name:"",
                email:"",
                description:""
            });
            list.select(nid);
            list.showItem(nid);
        },
        deleteUser : function(){
            var list = $$$("userlist");
            var id = list.getCursor();
            if (!id) return;
            webix.confirm({
                text:"URL <strong>\""+list.getItem(id).name+"\"</strong> 将被删除",
                callback:function(mode){
                    if (mode){
                        list.remove(id);
                        $$$("formpanel1").disable();
                        app.updateRelations(id);
                    }
                }
            });
        },
        toggleView:function(){
            var value = this.getValue();
            $$$(value).show();
        },
        deleteRole : function(){
            var grid = $$$("rolelist");
            var id = grid.getCursor();
            if (!id) return;
            webix.confirm({
                text:"<strong>\""+grid.getItem(id).name+"\"</strong> rule will be deleted",
                callback:function(mode){
                    grid.remove(id);
                    $$$("formpanel2").disable();
                    app.updateRelations(false, id);
                }
            });
        },
        AdRoleToUser : function(role){
            var user = $$$("userlist").getCursor();
            // if (user > 1000000000)
            // 	return webix.message({
            // 		type:"error",
            // 		text:"User's info contains errors, fix them first."
            // 	});
    
            // if (role > 1000000000)
            // 	return webix.message({
            // 		type:"error",
            // 		text:"Role's info contains errors, fix them first."
            // 	});
    
            var id = app.isUserHasRole(user, role);
            if (!id){
                app.relations.add({ user_id:user, role_id:role });
                //$$$("r2ulist").select(role, true);
                $$$("r2ulist").getItem(role).$check = true;
            } else {
                app.relations.remove(id);
                //$$$("r2ulist").unselect(role);
                $$$("r2ulist").getItem(role).$check = false;
            }
            $$$("r2ulist").refresh(role);
        },
        buildRelations:function(){
            var t = app.relationsHash = {};
            app.relations.data.each(function(obj){
                if (!t[obj.user_id])
                    t[obj.user_id] = [];
                t[obj.user_id].push(obj.role_id);
            });
        },
        isUserHasRole : function(user, role){
            var exists = false;
            app.relations.data.each(function(obj){
                if (obj.user_id == user && obj.role_id == role)
                    exists = obj.id;
            });
            return exists;
        },
        updateRelations:function(user, role){
            var ids = [];
            app.relations.data.each(function(obj){
                if (obj.user_id == user || obj.role_id == role)
                    ids.push(obj.id);
            });
    
            webix.dp(app.relations).ignore(function(){
                for (var i = 0; i < ids.length; i++)
                    app.relations.remove(ids[i]);	
            });
    
            app.buildRelations();
        },
        roleTemplate : "#name#<br><small>#description#</small>",
        checkTemplate : '<span class="webix_icon_btn fa-{obj.$check?check-:}square-o" style="max-width:32px;"></span>',
        closeButtonTemplate : "<span class='r2uclose'></span>"
    };
    
    app.relations = new webix.DataCollection();
    
    var rolelist = {
        view:"datatable", id:"rolelist", scroll:false,
        columns:[
            { id:"name", sort:"string", width:200, header:"Name" },
            { id:"description", header:"Comments", fillspace:true }
        ],
        select:"row"
    };
    var remove = function(){
        webix.confirm({
            text:"Are you sure", ok:"Yes", cancel:"Cancel",
            callback:function(res){
                if(res)return false}
        });
        return false
    };
    var form2 = {
        view:"toolbar",
        cols:[
          { view:"button", label:"Short", autowidth:true },
          { view:"button", label:"Medium", autowidth:true },
          { view:"button", label:"Long text", autowidth:true },
          { view:"button", label:"Very long text", autowidth:true },
          {}
        ]
      };
    var userlist ={ id:"userlist", view:"datatable", scroll:false,
        columns:[
            { id:"url", header:"URL", width:200, sort:"string" },
            { id:"pwd", header:"密码", width:200, sort:"string" },
            {id:"trash", header:"操作", align:"center",  template:'<button type="button" class="button_class">文件管理</button><button type="button" class="button_class">终端</button>',fillspace:true}
        ],
        onClick:{
			"webix_icon":remove
		},
        select:"row",
    };
    var userform = {
        type:"line", id:"formpanel1", rows:[
            { type:"header", css: "webix_header rounded_top", template:"添加URL" },
            { view:"form", id:"userform", width:400, elements:[
                { view:"text", label:"URL", name:"url" },
                { view:"text", label:"密码", name:"pwd" },
                // { view:"select", label:"选项", name:"option",value:'1',options:[
                //     { "id":1, "value":"自定义" },
                //     { "id":2, "value":"eval(暂未完成)" }
                // ]},
                { view:"textarea", label:"自定义配置", name:"description", height:345 },
                { view:"button",type: "form", label:"保存" , click: app.saveUserForm, inputWidth: 120, align: "center", height: 40},
                { height: 30 }
            ]}
        ]
    };
    var userview = { id:"user",height:"600", type:"wide", cols:[ userlist, userform ] };
    webix.protoUI({
        name:"readyuser",
        defaults:{
            type: "wide",
            rows:[
                { view:"toolbar", elements:[
                    { },
                    { view:"button", type:"icon", icon:"plus-circle", label:"Add", width:80, inputWidth:63,
                        click:app.addUser },
                    { view:"button", type:"icon", icon:"trash-o", label:"Delete", width:80,
                        click:app.deleteUser }
                ]},
                { animate: true, cells:[ userview ] }
            ]
        },
        $init:function(){
            var master = this;
            $$$ = function(name){
                return master.$$(name);
            }
    
            this.$ready.push(this._on_ui_created);
        },
        _on_ui_created:function(){
            webix.ajax(this.config.urls.data, function(text, data){
                data = data.json();
    
                $$$("userlist").parse(data.users);
                // $$$("rolelist").parse(data.roles);
                app.relations.parse(data.links);
                app.buildRelations();
    
                var first = $$$("userlist").getFirstId();
                if (first)
                    $$$("userlist").select(first);
            });
    
            $$$("userform").bind($$$("userlist"));
            webix.dp($$$("userlist")).define("url", this.config.urls.users);
    
            // $$$("roleform").bind($$$("rolelist"));
            // webix.dp($$$("rolelist")).define("url", this.config.urls.roles);
            webix.dp(app.relations).define("url", this.config.urls.links);
    
            $$$("formpanel1").disable();
            $$$("userlist").attachEvent("onAfterSelect", function(id){
                $$$("formpanel1").enable();
                // app.markUsedRoles(id.row);
                webix.delay(function(){
                    var input = $$$("userform").elements["url"].getInputNode();
                    input.select();
                    input.focus();
                });
            });
    
            // $$$("formpanel2").disable();
            // $$$("rolelist").attachEvent("onAfterSelect", function(id){
            //     $$$("formpanel2").enable();
            //     webix.delay(function(){
            //         var input = $$$("roleform").elements["name"].getInputNode();
            //         input.select();
            //         input.focus();
            //     });
            // });	
    
            // $$$('r2ulist').sync($$$("rolelist"));
            // $$$('r2ulist').attachEvent("onItemClick", function(id){
            //     app.addRoleToUser(id);
            //     app.buildRelations();
            // });
        }
    }, webix.IdSpace, webix.EventSystem, webix.ui.layout);
})();