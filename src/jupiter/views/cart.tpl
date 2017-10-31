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
            <h1>Your Purchase</h1>
            <!-- Item List Start -->
            <div class="item-label">
               <div class="item-name"><span>Item Name</span></div>
               <div class="item-quantity"><span>Quantity</span></div>
               <div class="item-price"><span>Price</span></div>
            </div>
            {{range .items}}
            <div class="item">
               <div class="item-name">{{.Name}}</div>
               <div class="item-quantity">{{.Qty}}</div>
               <div class="item-price">{{.Price}}</div>
            </div>
            {{end}}
            <!-- Item List End --> 
            <!-- NX Balance Start -->
            <div class="nx-balance">
               <div class="left-col">
                  <p>Select Payment Method:</p>
                  <div class="radio-row">
                     <input class="input-radio-settings" type="radio" name="nxgroup" value="3">
                     </input>
                     <label>NX Prepaid : {{.balance.Prepaid}}</label>
                  </div>
                  <div class="radio-row">
                     <input class="input-radio-settings" type="radio" name="nxgroup" value="2">
                     </input>
                     <label>NX Credit : {{.balance.Credit}}</label>
                  </div>
               </div>
               <div class="right-col">
                  <div class="refresh-button">Refresh</div>
                  <div class="row">
                     <div class="nx-balance-label">NX Prepaid</div>
                     <div class="nx-balance-number" id="nx-balance-selected"><strong>0</strong> NX</div>
                  </div>
                  <hr>
                  <div class="row">
                     <div class="nx-balance-label">Total Purchase Amount</div>
                     <div class="nx-balance-number" id="nx-balance-total"><strong>{{.invoice.TotalPrice}}.00</strong> NX</div>
                  </div>
                  <hr>
                  <div class="row">
                     <div class="nx-balance-label">Remaining NX Balance</div>
                     <div class="nx-balance-number" id="nx-balance-remained"><strong>0</strong> NX</div>
                  </div>
               </div>
               <div class="clear"></div>
            </div>
            <!-- NX Balance End -->
            <!-- Error Message Start -->
            <div class="error-message" style="display:none;">
               <p>[Error message] You have insufficient NX balance. <a class="link_purchase" href="/store/front/store/charge?ticket={{.ticket}}">Purchase NX</a></p>
            </div>
            <!-- Error Message End -->
            <p class="purchase-nx">Not enough NX? <a class="link_purchase" href="/store/front/store/purchase-nx?ticket={{.ticket}}">Purchase NX</a></p>
            <div class="button" id="btn_purcahse"><a id="link-place-order" href="{{.OrderURL}}">Place Order</a></div>
         </div>
<script type="text/javascript" charset="utf-8">
var bRequested = false;
var selected_nx = 0;
var order_url = "{{.OrderURL}}";
var nx_pp = 3;var nx_balace = { nx_pp:{{.balance.Prepaid}}, nx_cd:{{.balance.Credit}}, nx_total:{{.invoice.TotalPrice}} };
var nx_remained={{.balance.Total}};
function openPurchaseNXWin() {mbroker_publish("sdk-events",{"UIEvent":{"type":"OPEN_PURCHASE_NX_WINDOW", "url":"http://www.nexon.net/nx/purchase-nx/"}});}

function verifyNx() {
  $('.error-message').css({display:(nx_remained <=0) ? "block":"none"});
}
function fnRefresh() {
  bRequested = true;
  var newForm = $('<form>', {
      'action': '{{.refreshURL}}',
      'method' : 'get'
  }).append($('<input>', {
      'name': 'id',
      'ticket': '{{.ticket}}',
      'type': 'hidden'
  }));
  newForm.submit();
}
$(document).ready(function () {
  $('#btn_purcahse a').css({"text-decoration":"none"});
  $('.input-radio-settings').click(function () {
    nx_remained = ((this.value == nx_pp) ? nx_balace.nx_pp:nx_balace.nx_cd) - nx_balace.nx_total; 
    $('#nx-balance-selected').html("<strong>" + ((this.value == nx_pp) ? nx_balace.nx_pp:nx_balace.nx_cd).toFixed(2) +"</strong> NX"); 
    $('#nx-balance-remained').html("<strong>" + nx_remained.toFixed(2) +"</strong> NX"); 
	selected_nx = (this.value == nx_pp) ? nx_pp:2
	//alert(nx_remained)
    verifyNx();
  });
  $('#btn_purcahse a').click(function(e) {
    $('#btn_purcahse a').attr('href', order_url + "&cash_type=" + selected_nx.toString());
	//alert(nx_remained);
    if (nx_remained <=0) {
      e.preventDefault();
      verifyNx();
    }
    return true;
  });
  // open purchase window
  $('.link_purchase_').popupWindow({centerScreen:1,height:680,width:1070}); 
  $('.link_purchase_').click(function(e) {
    $('.modal-popup-container').css({display:"block"});
    openPurchaseNXWin();
    return false;
  });
  //dismiss modal
  $('#btn-modal-ok').click(function(e) {
    $('.modal-popup-container').css({display:"none"});
    // refresh this page
    // alert("clicked");
    //$('.modal-popup-container').nx_load_page('/shop/refreshCart','283348e6-2563-4544-938b-7cda6e150be7');
    var newForm = $('<form>', {
        'action': '/shop/refreshCart',
        'method' : 'post'
    }).append($('<input>', {
        'name': 'id',
        'value': '283348e6-2563-4544-938b-7cda6e150be7',
        'type': 'hidden'
    }));
    newForm.submit();
  });
  $('.input-radio-settings').first().trigger("click");
  $('.refresh-button').click(function(e) {
    if (bRequested) return false;
    fnRefresh();
    return false;
  });
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