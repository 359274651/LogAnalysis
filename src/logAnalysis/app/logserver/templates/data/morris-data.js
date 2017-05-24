

function refreshpie() {
    $.ajax({
        url:'/countstatusarea/1h',
        type:'GET', //GET
        async:true,    //或false,是否异步
        // data:{
        //     name:'yang',age:25
        // },
        // timeout:5000,    //超时时间
        dataType:'json',    //返回的数据格式：json/xml/html/script/jsonp/text
        // beforeSend:function(xhr){
        //     console.log(xhr)
        //     console.log('发送前')
        // },
        success:function(data,textStatus,jqXHR){
            console.log(data)
            console.log(textStatus)
            console.log(jqXHR)

            $.plot($("#flot-pie-chart"), data, {
                series: {
                    pie: {
                        show: true
                    }
                },
                grid: {
                    hoverable: true
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%p.0%, %s", // show percentages, rounding to 2 decimal places
                    shifts: {
                        x: 20,
                        y: 0
                    },
                    defaultTheme: false
                }
            });

        },
        error:function(xhr,textStatus){
            console.log('错误')
            console.log(xhr)
            console.log(textStatus)
        },
        complete:function(){
            console.log('结束')
        }
    })
}

//Flot Pie Chart
$(function() {

    refreshpie()

});






