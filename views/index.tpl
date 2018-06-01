<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>菜刀地址URL</title>
  <link rel="stylesheet" href="http://www.17sucai.com/preview/33522/2017-10-20/LarryCMS-new/common/layui/css/layui.css">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.css" media="all">
  <link rel="stylesheet" type="text/css" href="/static/css/global.css" media="all">
  <link rel="stylesheet" type="text/css" href="/static/css/personal.css" media="all">
</head>
<style>
  *{
    background: #fff;
  }
</style>
<body>
<section class="layui-larry-box">
    <div class="larry-personal">
      <blockquote class="layui-elem-quote news_search">
        <div class="layui-inline">
            <a class="layui-btn layui-btn-normal newsAdd_btn">添加</a>
        </div>
    </blockquote>
    <div class="layui-tab-content larry-personal-body clearfix mylog-info-box">
                 <!-- 操作日志 -->
                <div class="layui-tab-item layui-field-box layui-show">
                     <table class="layui-table" >
                          <thead>
                              <tr>
                                  <th><input type="checkbox" id="selected-all"></th>
                                  <th style="text-align:center;">ID</th>
                                  <th>URL</th>
                                  <th>密码</th>
                                  <th>操作</th>
                              </tr>
                          </thead>
                          <tbody id="caidao_list">
                          </tbody>
                     </table>
                </div>
        </div>
    </div>
</section>
<script type="text/javascript" src="static/layui/layui.js"></script>
<script type="text/javascript">
    function reflash(){
        location.reload();
    }
    layui.use(['jquery','layer','element','laypage'],function(){
          window.jQuery = window.$ = layui.jquery;
          window.layer = layui.layer;
          var element = layui.element;
          var laypage = layui.laypage;
          $.ajax({
              url :'/listjson',
              type:'get',
              success:function(msg){
                var msg = JSON.parse(msg);
                var str = '';
                  $.each(msg,function(index,result){
                      str += `
                            <tr>
                                <td><input type="checkbox"></td>
                                <td style="text-align:center;">${result.id}</td>
                                <td>${result.url}</td>
                                <td>${result.pwd}</td>
                                <td>
                                    <a target="_blank"  href="/wenjian?id=${result.id}" class="layui-btn layui-btn-radius layui-btn-normal">文件管理</a>
                                    <a target="_blank" data-url="${result.url}" data-id="${result.id}" class="layui-btn layui-btn-radius layui-btn-warm webssh">远程终端</a>
                                    <a target="_blank"  data-id="${result.id}" class="layui-btn layui-btn-radius layui-btn-danger delete">删除</a>
                                </td>
                            </tr>
                        `;
                  });
                  $("#caidao_list").html(str);
              }
          });
          $(document).on('click','.delete',function(){
              var id = $(this).attr('data-id');
              $.ajax({
                  url :"/deleteurl",
                  type:'post',
                  data:{id:id},
                  success:function(msg){}
              });
               $(this).closest('tr').remove();
          });
          $(document).on('click','.webssh',function(){
                var url  = $(this).attr('data-url');
                var array = url.split("/");
                var _url = array['0']+"\/\/"+array['2'] + "/conso.php";
                window.open(_url);  
          });
          $('.newsAdd_btn').click(function(){
              layer.open({
                title: '在线调试',
                type: 2,
                shadeClose: true,
                shade: false,
                maxmin: true, //开启最大化最小化按钮
                area: ['800px', '500px'],
                content: '/listadd',
              }); 
          });
    });
</script>
</body>
</html>