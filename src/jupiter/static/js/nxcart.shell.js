//
nxcart.shell = (function () {
	var
		configMap = {
			anchor_schema_map : {
				account_mapping : { open : true, clsoed : true },
				login : { open : true, clsoed : true },
				cmd : { status : ""}
			}
		},
		stateMap = { 
			$container : null,
			anchor_map : {},
			args : null
		},
		jqueryMap = {},
		setJqueryMap, setStep, setArgs, getArgs, toggleDialog, 
		onHashChange;
	//
	setJqueryMap = function() {
		var $container = stateMap.$container;
		setJqueryMap = { $container : $container};

	};
	setStep = function (step) {
		window.location.hash = "step=" + step;
	};
	setArgs = function (args) {
		stateMap.args = args;
		console.log('stateMap.args'+stateMap.args);
	};
	getArgs = function () {
		return stateMap.args;
	};
	toggleDialog = function ( status ) {
		//alert("toggle dialog : " + status);
		//$('#nxl-email-verification').hide();
		if (status == "open_charge") {
		    //$('.modal-popup-container').css({display:"block"});
			//$('#nxl-account-mapping').show();
		}
	};
	//
	onHashChange = function(event) {
		//alert("onHashChange");
		var
			anchor_map_proposed ;
		try {
			anchor_map_proposed = $.uriAnchor.makeAnchorMap();			
		}
		catch (error) {

		}
		stateMap.anchor_map = anchor_map_proposed;
		cmd = anchor_map_proposed.cmd;
		//alert("step : " + cmd);
		toggleDialog(cmd);
		return false;
	};
	initModule = function ( $container ) {
		//alert("called - initModule");
		$.uriAnchor.configModule ({
			schema_map : configMap.anchor_schema_map
		});
		$(window)
			.bind('hashchange', onHashChange)
			.trigger('hashchange');
			
		//$('#btnPlayGame').bind("click", onClickPlayGame);
	};
	return { 
		initModule : initModule, 
		setArgs : setArgs,
		getArgs : getArgs
	};
}());

//alert("nxcart.shell");
