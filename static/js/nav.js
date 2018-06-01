webix.protoUI({
	name:"navbar",
	defaults:{
		height:80,
		css:"navbar",
		template:"<div class='text'><a class='next' href='{obj.next.name}.html'></a><a class='prev' href='{obj.prev.name}.html'></a><span><a class='webix-btn btn-green btn-wide bigdevice' href='#link#'>Download Webix</a>#text#</span></div>"
	},

	value_setter:function(value){
		var index = 0;
		for (var i = 0; i < this.demos.length; i++)
			if (this.demos[i].name == value)
			index = i;
		this.data = this.demos[index];
		this.data.next = this.demos[(index+1) % this.demos.length];
		this.data.prev = this.demos[(index -1 + this.demos.length) % this.demos.length];
	},
	demos:[
		{ name:"pivot", 	link:"/pivot/download.html", text:"With Webix Pivot, you can easily create various complex table reports by dragging and dropping its elements graphically. It provides handy ways of summarizing, organizing, and comparing business data." },
		{ name:"kanban", 	link:"/kanban/download.html", text:"Create your personal tool for vizualization of workflow by using a flexible and easily embeddable Webix Kanban Board widget. You can customize its UI in accordance with project requirements." },
		{ name:"spreadsheet", 		link:"/spreadsheet/download.html", text:"Create your perfect spreadsheet with rich functionality using a flexible and easily embeddable Webix SpreadSheet widget.You can customize its UI in accordance with project requirements." },
		{ name:"filemanager",		link:"/filemanager/download.html", text:"Use Webix File Manager to build a customizable and easy-to-use tool for managing hierarchical data in your web apps. This widget allows dragging files and folders from one directory to another the way any desktop file explorer does." },
		
		{ name:"report",	link:"/download/", text:"Webix widgets help you create nice-looking sales reports. By combining both numerical and graphical data you can present sales and expenses per month, income and expense rates, year comparison, etc. as an interactive app." },
		{ name:"geo", 		link:"/download/", text:"It has never been so easy to create rich media apps. With Webix, you have a possibility to embed media content in the library widgets. Besides, the library supports integration with third-party tools that will let you add Google Maps to your apps." },
		{ name:"booking", 	link:"/download/", text:"Replace HTML form with Webix form and get all HTML5 goodies in all browsers, including Internet Explorer. Sliders, counters, data pickers and many other controls can be used as a part of Webix form." },
		{ name:"email", 	link:"/download/", text:"With Webix, data management becomes faster and more intuitive. This demo represents the main data components and the ways they can be used to show a hierarchical data." },
		{ name:"traders", 	link:"/download/", text:"Organizing and managing tabular data can be painless with Webix DataTable widget. Large amounts of related information can be put into easy-to-use tables which provide sorting, filtering and editing features." },
		{ name:"user", 		link:"/download/", text:"Instead of creating the same UI over and over again, try to create reusable components using Webix widgets. This complex demo is a single component, which can be reused in different applications." },
		{ name:"tracker", 	link:"/download/", text:"Are you willing to build responsive web apps? With Webix widgets, you have a possibility to develop apps that can sort, filter and group complex data on the client side at a high speed." },

		{ name:"shops", 	link:"/download/", text:"Organize your data navigation combining the functionality of Webix List and Accordion controls. Multiplied by the power of Webix DataTable, your ready application will be able to boast of ultimate user experience." },
		{ name:"touch", 	link:"/download/", text:"You can develop complex and trendy mobile web apps that will include good-looking charts, long lists with kinetic scrolling,  and tabbars adapted for mobile screens and  sliders. All desktop features can be used in your mobile apps." },
		{ name:"scheduler", link:"/scheduler/download.html", text:"With Webix Scheduler, you can create interactive event calendars optimized for all types of mobile devices. This JavaScript calendar supports three basic views (Day, Month, Year), as well as single-, multi-day and recurring events." },
		{ name:"pivot_chart", 	link:"/pivot/download.html", text:"With Webix Pivot Chart you can easily create a graphical representation of business data. It provides handy ways of organizing, viewing and comparing complex data sets." }
	]
}, webix.ui.template );