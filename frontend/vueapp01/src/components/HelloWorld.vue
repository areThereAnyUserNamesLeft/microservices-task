<template>
    <div class="hello">
        <p>Click the user name for more details:</p>
        <ul>
            <li v-for="(n,index) in ip" v-on:click="getUserDetails(index + 1)" >
                {{ ip[index].user_name }} is currently "{{ip[index].user_status}}" - <span class="onhover" v-on:click="deleteUserDetails(index + 1)">DELETE</span>
            </li>
        </ul>
        <br />
        <br />
        <textarea>{{ response }}</textarea>
        <p>Add users (Press F5 to refresh, if not instantly added):</p>
        <input type="text" v-model="input.user_name" placeholder="User Name" />
        <input type="text" v-model="input.user_status" placeholder="User Status" />
        <input type="date" v-model="input.date" />
        <button v-on:click="sendData()">Add</button>

    </div>
</template>

<script>
    import axios from "axios";

export default {
        name: 'HelloWorld',
        data () {
            return {
                ip: "",
                input: {
                    user_name: "",
                    user_status: "",
                    date: ""
                },
                response: ""
            }
        },
        mounted() {

        axios({ method: "GET", "url": "http://127.0.0.1:8090/api/v1/users"}).then(result => {
            this.ip = result.data
        }, error =>{
            console.error(error);
        });
    },
    methods: {
            sendData() {
                axios({ method: "POST", "url": "http://127.0.0.1:8090/api/v1/users", "data":this.input, "headers": {"content-type": "application/json" } }).then(result => {
                    this.response = result.data;
                }, error => {
                    console.error(error);
                });


            },
            getUserDetails(i) {
                axios({ method: "GET", "url":"http://127.0.0.1:8090/api/v1/users/" + i, "data":this.input, "headers": {"content-type": "application/json" } }).then(result => {
                    this.response = result.data;
                }, error => {
                    console.error(error);
                });


            },
            deleteUserDetails(i) {
                axios({ method: "DELETE", "url":"http://127.0.0.1:8090/api/v1/users/" + i, "data":this.input, "headers": {"content-type": "application/json" } }).then(result => {
                    this.response = result.data;
                }, error => {
                    console.error(error);
                });


            }
    }
}
</script>

<style scoped>

@import url('https://fonts.googleapis.com/css?family=Cabin+Sketch|VT323');

    h1, h2 {
        font-weight: normal;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        font-family: 'Cabin Sketch', Cursive;
        font-size: 30px;
        display: block;
        margin: 0 10px;
        padding-bottom: 10px;
    }

    a {
        color: #42b983;
    }
::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
    color: whitesmoke;
    opacity: 1; /* Firefox */
}
    input[type="text"], input[type="date"], textarea {
        font-family: 'VT323', monospace;
        font-size: 25px;
        color: whitesmoke;
        background-color: grey;
        border: none;
    }
    input[type="text"]{
        padding: 10px;
    }
    input[type="date"]{
        padding: 4px;
    }
    textarea {
        width: 600px;
        height: 200px;
    }
    button {
        border: none;
        background: green;
        font-family: 'Cabin Sketch', Cursive;
        padding: 15px;
        color: whitesmoke;
    }
    .onhover {
        color: red;
        display: none;
    }
    li:hover .onhover {
        display:inline;
    }
</style>
