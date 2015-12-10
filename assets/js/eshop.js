$(document).ready(function(){
	$("#menu-item-address").click(function(){
		$("#address-modal").modal('show');
	});
	$("#address-new-btn").click(function(){
		$("#address-modal").modal('hide');
		$("#address-form-modal").modal('show');
	});
	$("#address-form-submit").click(function(){
		$("#address-form").submit();
	});
	$(".address-del-btn").click(function(){
		var id = $(this).data("id");
		$.post("/address/del",{id : id},function(result){
			 window.location.replace("/")
		});
	});
	$(".goods_buy_btn").click(function(){
		$("#order-goods-img").attr("href", "/assets/data/goods/" + $(this).data("img"));
		$("#order-goods-name").text($(this).data("name") + "        ￥"+$(this).data("price"));
		$("#order-form-modal").modal('show');
	});
});