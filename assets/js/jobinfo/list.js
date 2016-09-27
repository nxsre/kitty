/**
 * Created by dongtian on 16/9/27.
 */

var JobInfoList = {


    init:function () {

    this.ctrl.initEvent();

    },
    ctrl:{
        initEvent:function () {

            $("[name='Active']").bootstrapSwitch({
                onSwitchChange: function (e, state) {
                    alert(e);
                    alert(state);
                    return false;
                }
            });

            $('.btn_delJob').bind('click',function () {

                layer.msg('你确定删除此任务吗？', {
                    time: 0 //不自动关闭
                    ,btn: ['确定', '取消']
                    , icon: 6
                    ,yes: function(index){
                        layer.close(index);

                    }
                });

            });


        }
    }
};



$(function () {

    JobInfoList.init();
});