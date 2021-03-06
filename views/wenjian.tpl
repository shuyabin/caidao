<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    <link rel="stylesheet" href="/static/webix/webix.css">
    <link rel="stylesheet" href="/static/webix/filemanager.css" type="text/css" charset="utf-8">
</head>
<body>
    <div id="myTabContent" class="tab-content">
            <div id="my_box">
            </div>
    </div>
</body>
<script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
<script src="/static/webix/webix.js"></script>
<script src="/static/webix/filemanager.js" type="text/javascript"></script>
<script>
	function GetRequest() {   
   var url = location.search; //获取url中"?"符后的字串   
   var theRequest = new Object();   
   if (url.indexOf("?") != -1) {   
      var str = url.substr(1);   
      strs = str.split("&");   
      for(var i = 0; i < strs.length; i ++) {   
         theRequest[strs[i].split("=")[0]]=unescape(strs[i].split("=")[1]);   
      }   
   }   
   return theRequest;   
}   
	var res = GetRequest();	
	var id = res.id;
	webix.ready(function(){
		webix.i18n.filemanager = {
	    actions: "Actions",
	    back: "Back",
	    forward: "Forward",
	    levelUp: "Level Up",
	    name: "Name",
	    size: "大小",
	    type: "类型",
	    date: "时间",
	    copy: "复制",
	    cut: "剪切",
	    paste: "粘贴",
	    upload: "上传",
	    remove: "删除",
	    create: "创建文件夹",
	    rename: "重命名",
	    location: "Location",
	    select: "Select Files",
	    sizeLabels: ["B","KB","MB","GB"],
	    iconsView: "Icons View",
	    tableView: "Table View",
	    hideTree: "Hide Tree",
	    showTree: "Show Tree",
	    collapseTree: "Collapse Tree",
	    expandTree: "Expand Tree",
	    saving: "Saving...",
	    errorResponse: "Error: changes were not saved!",
	    replaceConfirmation: "Would you like to replace existing files?",
	    createConfirmation: "The folder already exists. Would you like to replace it?",
	    renameConfirmation: "The file already exists. Would you like to replace it?",
	    yes: "Yes",
	    no: "No",
	    types:{
	        folder: "文件夹",
	        doc: "文档",
	        excel: "Excel",
	        pdf: "PDF",
	        pp: "PowerPoint",
	        text: "文本文件",
	        video: "视频文件",
	        image: "图片",
	        code: "Code",
	        audio: "Audio",
	        archive: "Archive",
	        file: "File"
	    }
	};
    //object constructor
    webix.ui({
        view:"filemanager",
        id: "files", height: 700, container: "my_box",
        handlers: {
            "files": "filelist?d_id="+id,
            "search": "search",
            "upload": "fileupload?d_id="+id,
            "download": "download?d_id="+id,
            "copy": "copy?d_id="+id,
            "move": "move?d_id="+id,
            "remove": "delete?d_id="+id,
            "rename": "rename?d_id="+id,
            "create": "newFolder?d_id="+id
        }
    });
    // loading data source
    $$("files").load('/filelist?d_id='+id);
    $$("files").attachEvent("onBeforeRun", function (id) {
        webix.confirm({
            text: "您要下载本文件么?",
            ok: "下载",
            cancel: "取消",
            callback: function (result) {
                if (result)
                    $$("files").download(id);
            }
        });
        return false;
    });
});
</script>
</html>