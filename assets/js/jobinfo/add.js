/**
 * 新建任务
 * Created by dongtian on 16/9/20.
 */



var JobInfo = {

    init: function () {
        this.ctrl.initEvent();
        this.ctrl.initValidation();
    },
    ctrl: {

        initEvent: function () {
            $('#btnSave').bind('click', function () {

                var formvalidation = $('#addFrom').data('formValidation');
                formvalidation.validate();
                var valid = formvalidation.isValid();
                if (valid == false) {
                    return false;
                }

                JobInfo.ctrl.addJob();
                return false;
            });
            
            $('#btn_editJob').bind('click',function () {


            });
        },

        addJob: function () {
            var formDat = $('#addFrom').serialize();
            layer.load(2);
            $.ajax({
                url: '/jobinfo/add',
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

        },

        initValidation: function () {

            $('#addFrom').formValidation({
                framework: 'bootstrap',
                icon: {
                    valid: 'glyphicon glyphicon-ok',
                    invalid: 'glyphicon glyphicon-remove',
                    validating: 'glyphicon glyphicon-refresh'
                },
                fields: {
                    JobName: {
                        validators: {
                            notEmpty: {
                                message: '任务名称不能为空!'
                            },
                            stringLength: {
                                min: 3,
                                max: 30,
                                message: '任务名称必须介于3至32字符!'
                            },
                            regexp: {
                                regexp: /^[a-zA-Z0-9_\u4E00-\u9FA5]+$/,
                                message: '任务名称必须中文、字母、数字、_ 之间!'
                            }
                        }
                    },
                    JobGroup: {
                        validators: {
                            notEmpty: {
                                message: '任务分组不能为空!'
                            },
                            stringLength: {
                                min: 3,
                                max: 30,
                                message: '任务分组必须介于3至32字符!'
                            },
                            regexp: {
                                regexp: /^[a-zA-Z0-9_\u4E00-\u9FA5]+$/,
                                message: '任务分组必须中文、字母、数字、_之间!'
                            }
                        }
                    },
                    Url: {
                        validators: {
                            notEmpty: {
                                message: '目标服务器地址Url不能为空!'
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
        }
    }
};

$(function () {

    JobInfo.init();

});