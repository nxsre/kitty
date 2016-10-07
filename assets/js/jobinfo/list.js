/**
 * Created by dongtian on 16/9/27.
 */

var JobInfoList = {


    init: function () {

        this.ctrl.initEvent();

    },
    ctrl: {
        initEvent: function () {

            $("[name='Active']").bootstrapSwitch({
                onSwitchChange: function (e, state) {
                    alert(e);
                    alert(state);
                    return false;
                }
            });

            $('.btn_delJob').bind('click', function () {

                var id = $(this).attr('att');
                layer.msg('你确定删除此任务吗？', {
                    time: 0 //不自动关闭
                    , btn: ['确定', '取消']
                    , icon: 6
                    , yes: function (index) {
                        layer.close(index);
                        JobInfoList.ctrl.deleteJobInfo(id);
                    }
                });

            });


        },
        deleteJobInfo: function (id) {

            var formDat = {"Id": id};
            layer.load(2);
            $.ajax({
                url: '/jobinfo/delete',
                cache: false,
                data: formDat,
                dataType: 'json',
                type: 'POST',

                error: function (req, status, err) {
                    layer.closeAll('loading');
                    layer.alert('提交失败,请重试!', {icon: 5});
                },
                success: function (data) {
                    layer.closeAll('loading');
                    if (data.success == true) {

                        layer.msg(data.message, function () {
                            window.location.href = "/jobinfo/list";
                        });

                    } else {
                        layer.msg(data.message, {icon: 5});
                    }
                }

            });
        }
    }
};


$(function () {

    JobInfoList.init();
});