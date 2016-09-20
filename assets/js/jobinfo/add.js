/**
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
                var valid = formvalidation.isValid();
                if (valid == true) {
                    alert("true");
                } else {
                    alert("false");
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
                                message: '任务名称必须介于3~32字符之间!'
                            },
                            regexp: {
                                regexp: /^[a-zA-Z0-9_]+$/,
                                message: '任务名称必须字母、数字、_ 之间!'
                            }
                        }
                    },
                    password: {
                        validators: {
                            notEmpty: {
                                message: 'The password is required'
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