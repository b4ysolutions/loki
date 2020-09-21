<template>
  <div class="main">
    <!-- START Loki Info -->
    <b-container class="loki-info">
      <img alt="Vue logo" src="../assets/logo.png" />
      <h1 class="top-title">{{ msg }}</h1>

      <b-container class="loki-short-description">
        <p>
          Loki webservice was created with the purpose of helping with automated
          tests.
        </p>
        <p>
          It arose from the need to have an endpoint to validate which load
          testing framework would be the most suitable among the possibilities
          presented. As nothing similar and free was found, Loki was developed
          to be a test endpoint solution for everyone.
        </p>
      </b-container>
    </b-container>
    <!-- START Loki Info -->

    <!-- START Getting Started -->
    <h2 class="main-title bg-light" id="getting-started">Getting Started</h2>
    <b-container>
      <h3 id="repository">Repository</h3>
      <p>
        The current version of the Loki API lives at
        <a href="https://github.com/b4ysolutions/loki-api">github.com/b4ysolutions/loki-api</a>.
      </p>

      <h3 id="versions">Versions</h3>
      <b-table small :items="items"></b-table>

      <h3 id="requirements">Requirements</h3>
      <p>
        <span class="requirements">Go</span> -
        <a target="_Blank" href="https://golang.org/"
          ><open-in-new-icon class="open-in-new-icon"
        /></a>
      </p>
      <p>
        <span class="requirements">MySQL</span>-
        <a target="_Blank" href="https://www.mysql.com/"
          ><open-in-new-icon class="open-in-new-icon"
        /></a>
      </p>

      <h3 id="usage">Usage</h3>
      <p>
        Download and unzip the lastest version:
        <BtnCopy :text="usage_steps[0]" />
      </p>
      <div class="code">{{ usage_steps[0] }}</div>
      <p>Go to the project folder: <BtnCopy :text="usage_steps[1]" /></p>
      <div class="code">{{ usage_steps[1] }}</div>
      <p>Start Loki webservice: <BtnCopy :text="usage_steps[2]" /></p>
      <div class="code">{{ usage_steps[2] }}</div>
    </b-container>
    <!-- END Getting Started -->

    <!-- START Endpoit Users -->
    <h2 class="main-title bg-light" id="endpoit-users">USERS</h2>
    <Endpoint
      id="endpoint-users-list"
      title="List"
      description="List all registered users."
      url="http://localhost:8787/users"
      :examples="[
        {
          id: 'users-list-1',
          url: 'http://localhost:8787/users',
          method: 'GET',
          title: 'Example of List with Users registered',
          body: {},
          response: {
            id: 2003,
            type: 'success',
            message: 'List with all users',
            users: [
              {
                id: 1,
                name: 'CreateTest',
                email: 'create@test.com',
                password: 'createtest',
                status: 1,
                islogged: true,
              },
              {
                id: 2,
                name: 'CreateTest2',
                email: 'create@test2.com',
                password: 'createtest2',
                status: 3,
                islogged: false,
              },
            ],
          },
        },
        {
          id: 'users-list-2',
          url: 'http://localhost:8787/users',
          method: 'GET',
          title: 'Example of List with no Users registered',
          body: {},
          response: {
            id: 2008,
            type: 'success',
            message: 'No user found',
          },
        },
      ]"
    />
    <Endpoint
      id="endpoint-users-get"
      title="Get"
      description="Get a registered user."
      url="http://localhost:8787/users/:id"
      :examples="[
        {
          id: 'users-get-1',
          url: 'http://localhost:8787/users/1',
          method: 'GET',
          title: 'Example of Get with User registered',
          body: {},
          response: {
            id: 2004,
            type: 'success',
            message: 'User found',
            users: [
              {
                id: 1,
                name: 'CreateTest',
                email: 'create@test.com',
                password: 'createtest',
                status: 3,
                islogged: false,
              },
            ],
          },
        },
        {
          id: 'users-get-2',
          url: 'http://localhost:8787/users/1',
          method: 'GET',
          title: 'Example of Get with no User registered',
          body: {},
          response: {
            id: 2009,
            type: 'success',
            message: 'No user found',
          },
        },
      ]"
    />
    <Endpoint
      id="endpoint-users-create"
      title="Create"
      description="Create a new user."
      url="http://localhost:8787/users"
      :examples="[
        {
          id: 'users-create-1',
          url: 'http://localhost:8787/users',
          method: 'POST',
          title: 'Example of Create using a new email',
          body: {
            name: 'CreateTest',
            email: 'create@test.com',
            password: 'createtest',
            status: 1,
            islogged: false,
          },
          response: {
            id: 2005,
            type: 'success',
            message: 'User successfully created',
            users: [
              {
                id: 1,
                name: 'CreateTest',
                email: 'create@test.com',
                password: 'createtest',
                status: 1,
                islogged: false,
              },
            ],
          },
        },
        {
          id: 'users-create-2',
          url: 'http://localhost:8787/users',
          method: 'POST',
          title: 'Example of Create using an existing email',
          body: {
            name: 'CreateTest',
            email: 'create@test.com',
            password: 'createtest',
            status: 1,
            islogged: false,
          },
          response: {
            id: 1006,
            type: 'error',
            message: 'Failed to create the user',
          },
        },
      ]"
    />
    <Endpoint
      id="endpoint-users-update"
      title="Update"
      description="Update a registered user."
      url="http://localhost:8787/users"
      :examples="[
        {
          id: 'users-update-1',
          url: 'http://localhost:8787/users',
          method: 'PUT',
          title: 'Example of Update using a registered user',
          body: {
            id: 1,
            name: 'UpdateTest',
            email: 'update@test.com',
            password: 'updatetest',
            status: 3,
            islogged: true,
          },
          response: {
            id: 2002,
            type: 'success',
            message: 'User successfully updated',
            users: [
              {
                id: 1,
                name: 'UpdateTest',
                email: 'update@test.com',
                password: 'updatetest',
                status: 3,
                islogged: true,
              },
            ],
          },
        },
        {
          id: 'users-update-2',
          url: 'http://localhost:8787/users',
          method: 'PUT',
          title: 'Example of Update using a not registered user',
          body: {
            id: 87,
            name: 'UpdateTest',
            email: 'update@test.com',
            password: 'updatetest',
            status: 3,
            islogged: true,
          },
          response: {
            id: 2007,
            type: 'success',
            message: 'User not updated',
          },
        },
      ]"
    />
    <Endpoint
      id="endpoint-users-delete"
      title="Delete"
      description="Delete a registered user."
      url="http://localhost:8787/users/:id"
      :examples="[
        {
          id: 'users-delete-1',
          url: 'http://localhost:8787/users/1',
          method: 'DELETE',
          title: 'Example of Delete using a registered user',
          body: {},
          response: {
            id: 2001,
            type: 'success',
            message: 'User successfully removed',
          },
        },
        {
          id: 'users-delete-2',
          url: 'http://localhost:8787/users/2',
          method: 'DELETE',
          title: 'Example of Delete using a not registered user',
          body: {},
          response: {
            id: 2006,
            type: 'success',
            message: 'User not removed',
          },
        },
      ]"
    />
    <Endpoint
      id="endpoint-users-login"
      title="Login"
      description="Login to a registered user."
      url="http://localhost:8787/users/login"
      :examples="[
        {
          id: 'users-login-1',
          url: 'http://localhost:8787/users/login',
          method: 'PUT',
          title: 'Example of Login using a not logged user',
          body: {
            email: 'create@test.com',
            password: 'createtest',
          },
          response: {
            id: 2010,
            type: 'success',
            message: 'User successfully login',
            users: [
              {
                id: 1,
                name: 'CreateTest',
                email: 'create@test.com',
                password: 'createtest',
                status: 1,
                islogged: true,
              },
            ],
          },
        },
        {
          id: 'users-login-2',
          url: 'http://localhost:8787/users/login',
          method: 'PUT',
          title: 'Example of Login using a logged user',
          body: {
            email: 'create@test.com',
            password: 'createtest',
          },
          response: {
            id: 2011,
            type: 'success',
            message: 'User is already logged',
            users: [
              {
                id: 1,
                name: 'CreateTest',
                email: 'create@test.com',
                password: 'createtest',
                status: 1,
                islogged: true,
              },
            ],
          },
        },
      ]"
    />
    <Endpoint
      id="endpoint-users-logout"
      title="Logout"
      description="Logout to a registered user."
      url="http://localhost:8787/users/logout/:id"
      :examples="[
        {
          id: 'users-logout-1',
          url: 'http://localhost:8787/users/logout/1',
          method: 'GET',
          title: 'Example of Logout using a logged user',
          body: {},
          response: {
            id: 2012,
            type: 'success',
            message: 'User successfully logout',
            users: [
              {
                id: 1,
                name: 'CreateTest',
                email: 'create@test.com',
                password: 'createtest',
                status: 1,
                islogged: false,
              },
            ],
          },
        },
        {
          id: 'users-logout-2',
          url: 'http://localhost:8787/users/logout/1',
          method: 'GET',
          title: 'Example of Login using a not logged user',
          body: {},
          response: {
            id: 2013,
            type: 'success',
            message: 'User is not logged',
          },
        },
      ]"
    />
    <!-- END Endpoit Users -->

    <div class="footer">
      <a href="https://b4ysolutions.com">
        <img alt="Vue logo" src="../assets/b4y_big.png" class="logo-b4y-big" />
      </a>
    </div>
  </div>
</template>

<script>
import Endpoint from "./Endpoint.vue";
import BtnCopy from "./BtnCopy.vue";
import OpenInNewIcon from "vue-material-design-icons/OpenInNew.vue";

export default {
  name: "Main",
  props: {
    msg: String,
  },
  components: {
    Endpoint,
    BtnCopy,
    OpenInNewIcon,
  },
  data() {
    return {
      items: [
        {
          version: "1.0.0",
          date: "01/09/2020",
          main_changes: "Initial deployment",
        },
      ],
      usage_steps: [
        "tar -xvf loki-api-1.0.0.tar.gz",
        "cd loki-api-1.0.0",
        "go run main.go",
      ],
    };
  },
};
</script>

<style lang="scss" scoped>
.open-in-new-icon {
  color: var(--color-fun-green);
}
span.requirements {
  padding-left: 10px;
  font-style: italic;
  font-weight: bold;
}
.footer {
  width: 100%;
  text-align: center;  
  background: var(--color-wild-sand);
  padding: 50px 0 20px 0;
  margin-top: 100px;

  .logo-b4y-big {
    width: auto;
    margin: 5px 0;
  }
}
</style>
