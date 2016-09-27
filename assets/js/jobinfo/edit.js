/**
 * Created by dongtian on 16/9/27.
 */

var EditJobInfo = {
    init: function () {
        this.ctrl.initFormValidation();
        this.ctrl.initEvent();
    },
    ctrl: {

        initEvent: function () {
            
            $('#btnSave').bind('click',function () {

                var formvalidation = $('#editForm').data('formValidation');
                formvalidation.validate();
                var valid = formvalidation.isValid();
                if (valid == false) {
                    return false;
                }

                EditJobInfo.ctrl.editJob();
                return false;
            });

            $('input').iCheck({
                checkboxClass: 'icheckbox_flat-blue',
                radioClass: 'iradio_flat-blue'
            });
            $('input').iCheck('disable');
        },
        initFormValidation:function () {

            $('#editForm').formValidation({
                framework: 'bootstrap',
                icon: {
                    valid: 'glyphicon glyphicon-ok',
                    invalid: 'glyphicon glyphicon-remove',
                    validating: 'glyphicon glyphicon-refresh'
                },
                fields: {
                    Url: {
                        validators: {
                            notEmpty: {
                                message: '目标服务器地址Url不能为空!'
                            },
                            regexp: {
                                regexp: /^(http|https):\/\/(\d+\.){3}(\d+)(:\d+)?(\S)*$/,
                                message: '目标服务器地址Url格式不正确!'
                            }
                        }
                    },
                    Cron: {
                        validators: {
                            notEmpty: {
                                message: 'Cron表达式不能为空!'
                            }
                        }
                    },
                    Params: {
                        validators: {
                            notEmpty: {
                                message: '目标服务器地址参数不能为空!'
                            }
                        }
                    }
                }
            });

        },
        editJob:function () {


            var formDat = $('#editForm').serialize();
            layer.load(2);
            $.ajax({
                url: '/jobinfo/edit',
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

                        layer.msg('保存成功', function () {
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

    EditJobInfo.init();
});
