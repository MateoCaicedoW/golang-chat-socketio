const socket = io();
const s2 = io("/chat");
new Vue({
  el: '#chat-app',
  created() {
    s2.on("reply", (data) => {
      this.messages.push({
        text: data,
        date: new Date().toLocaleString()
      })
    })
  },
  data: {
    message: '',
    messages: []
  },
  methods: {
    sendMessage() {
      s2.emit("msg", this.message)
      console.log(this.message);
      this.message = "";
    }
  }
})


