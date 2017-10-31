<!DOCTYPE html>
<html lang="en">
   <head>
      <meta http-equiv="Content-Type" content="text/html;charset=UTF-8" />
      <title>Purchase Confirmation</title>
      <link rel="stylesheet" type="text/css" href="//s3.amazonaws.com/nxcache/nxl/css/purchase_confirm.css" />
      <meta name="description" content="jiojio" />
      <link rel="shortcut icon" type="image/x-ico" href="//s3.amazonaws.com/nxcache/all/favicon.ico" />
      <meta name="viewport" content="width=device-width" />
      <script src="//s3.amazonaws.com/nxcache/nxl/js/jquery-1.7.1.min.js" type="text/javascript"></script>
      <script src="//s3.amazonaws.com/nxcache/nxl/js/jquery.popupWindow.js" type="text/javascript"></script>
      <script src="//s3.amazonaws.com/nxcache/nxl/js/jquery.nexon.shop.ui-0.1.js" type="text/javascript"></script>
      <!-- test -->
      <!-- script src="//s3.amazonaws.com/nxcache/nxl/js/mbroker.js" type="text/javascript"></script  -->
   </head>
   <body>
      <div class="wrap">
         <div id="nxcart" class="content">
            <!-- contents Start -->
            <h1>Select Item...</h1>
            <!-- Item List Start -->
            <div class="item-label">
               <div class="item-name"><span>Item Name</span></div>
               <div class="item-quantity"><span>Quantity</span></div>
               <div class="item-price"><span>Price</span></div>
            </div>
            <form id="frm" action="/store/demo/cart?ticket={{.ticket}}" method="get">
			  <input type="hidden" name="ticket" value="{{.ticket}}">
               {{range .items}}
               <div class="item">
                  <div class="item-name"><input type="checkbox" name="item" value="{{.Id}}">&nbsp;{{.Name}}</div>
                  <div class="item-quantity">{{.Qty}}</div>
                  <div class="item-price">{{.Price}}</div>
               </div>
               {{end}}
            </form>
            <!-- Item List End --> 
            <!-- Error Message Start -->
            <div class="error-message" style="display:none;">
               <p>[Error message] You have insufficient NX balance. <a class="link_purchase" href="/store/front/store/charge?ticket={{.ticket}}">Purchase NX</a></p>
            </div>
            <!-- Error Message End -->
            <div class="button" id="btn_purcahse"><a id="link-place-order" >Check Out</a></div>
         </div>
<script type="text/javascript" charset="utf-8">
var bRequested = false;
var selected_nx = 0;
var order_url = "{{.OrderURL}}";
$('#btn_purcahse a').click(function(e) {
	$('#frm').submit();			
});

</script>
      </div>
      <!-- Popup Start -->
      <div class="modal-popup-container" style="display:none;">
         <div class="modal-popup">
            <div class="modal-popup-content">
               <p><strong>Purchasing NX...</strong></p>
               <p>Click Ok if you are done with purchasing NX.</p>
               <div class="flush-bottom">
                  <hr>
                  <div class="button-small" id="btn-modal-ok">Ok</div>
               </div>
            </div>
         </div>
      </div>
   </body>
</html>
