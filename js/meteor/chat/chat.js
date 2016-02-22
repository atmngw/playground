var Messages = new Meteor.Collection('messages');

if (Meteor.isClient) {
    Template.contents.messages = function () {
        return Messages.find({}, {sort:{date:-1}} );
    }

    Template.contents.events({
        'keydown input#postmes':function(ev){
            if(ev.keyCode == 13){ // enter key
                var message = $("#postmes").val();
                if(message != ""){
                    $("#postmes").val("");
                    var setdata = createProperties(message);
                    // コレクションへ新レコードを登録
                    Messages.insert(setdata, function(err, _id){
                        // 最大件数以上なら古いのを消す
                        var len = Messages.find({}).fetch().length;
                        if(len > maxrec){
                            var doc = Messages.findOne({}); // 先頭レコード(つまり一番古いレコード)を取り出す
                            Messages.remove({_id:doc._id}); // そして消す
                        }
                    })
                }
            }
        }
    });

    // messageレコードのプロパティ生成
    function createProperties(message){
        var date = Date.parse(new Date());
        var datetime = toDateStr(date);
        var name = $("#postname").val();
        var style = color;
        if(name == ""){
            name = "guest";
        }
        return {date:date,datetime:datetime,message:message,name:name,style:style};
    }

    // 日付を文字列を変換
    function toDateStr(parseDate){
        var date = new Date(parseDate);
        var y = date.getFullYear();
        var m = date.getMonth()+1;
        var d = date.getDate();
        var h = date.getHours();
        var min = date.getMinutes();
        return y+"/"+m+"/"+d+" "+h+":"+min;
    }

    var maxrec = 10;
    Template.contents.maxrec = maxrec; // 最大表示件数
    var colors = [
      'Cornsilk','Aquamarine','LightPink','PaleGreen','Coral',
      'SpringGreen','Plum','LightSkyBlue','MistyRose','Turquoise'];
    var color = "background-color:"+colors[(+new Date()) % 10]+";";
    Template.contents.color = color;    // 背景色表示
}

if (Meteor.isServer) {
  Meteor.startup(function () {
      Messages.remove({});
  });
}
