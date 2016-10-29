/**
 * Created by dongtian on 16/9/27.
 */

var JobInfoList = {


    init: function () {

        this.ctrl.initEvent();

    },
    ctrl: {
        initEvent: function () {
            $("[name='Active']").bootstrapSwitch();

            $("[name='Active']").on('switchChange.bootstrapSwitch', function (event, state) {
                $(this).bootstrapSwitch('state', !state, true);
                var id = $(this).val();
                JobInfoList.ctrl.activeJob(id, state, $(this));

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

            //
            $('#btn_submit').click(function () {
                $('#pageSize').val(10);
                $('#pageNo').val(1);
                $('#searchForm').submit();

            });

        },

        // 激活
        activeJob: function (id, active, ele) {

            var msg = "你确定要激活此任务吗";
            var activeV = 1;
            if (active == false) {
                var msg = "你确定要取消此任务吗";
                activeV = 0;
            }
            layer.msg(msg, {
                time: 0 //不自动关闭
                , btn: ['确定', '取消']
                , icon: 6
                , yes: function (index) {
                    layer.close(index);
                    var formDat = {"id": id, "active": activeV};

                    $.ajax({
                        url: '/jobinfo/active',
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

                                if (active == true) {
                                    ele.bootstrapSwitch('state', true, true);
                                } else {
                                    ele.bootstrapSwitch('state', false, true);
                                }

                            } else {
                                layer.msg(data.message, {icon: 5});
                            }
                        }

                    });


                }
            });
        },


        pageFun: function (pageSize, pageNo) {

            $('#pageSize').val(pageSize);
            $('#pageNo').val(pageNo);
            $('#searchForm').submit();

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