<template>
  <div>
    <b-navbar fixed="top" type="dark" variant="light">
      <b-container>
        <h1 class="page-title-left" v-scroll-to="'#app'">
          <img alt="Vue logo" src="../assets/logo.png" class="logo-min" />Loki
        </h1>

        <b-navbar-nav class="ml-auto">
          <!-- START - Left Menu -->
          <div>
            <a v-b-toggle href="#sidebar" class="global-icon" @click.prevent>
              <menu-icon class="global-icon" />
            </a>
            <b-sidebar id="sidebar" title shadow>
              <div class="px-0 py-2 menu-left">
                <div>
                  <b-nav vertical class="w-100 left-menu">
                    <b-nav-item
                      ><p v-scroll-to="'#getting-started'" class="menu-item">
                        Getting Started
                      </p>
                      <b-nav vertical class="w-100 left-menu">
                        <b-nav-item
                          ><p v-scroll-to="'#repository'" class="submenu-item">
                            <chevron-right-icon />Repository
                          </p></b-nav-item
                        >
                        <b-nav-item
                          ><p v-scroll-to="'#versions'" class="submenu-item">
                            <chevron-right-icon />Versions
                          </p></b-nav-item
                        >
                        <b-nav-item
                          ><p v-scroll-to="'#requirements'" class="submenu-item">
                            <chevron-right-icon />Requirements
                          </p></b-nav-item
                        >
                        <b-nav-item
                          ><p v-scroll-to="'#usage'" class="submenu-item">
                            <chevron-right-icon />Usage
                          </p></b-nav-item
                        >
                      </b-nav>
                    </b-nav-item>
                    <b-nav-item
                      ><p v-scroll-to="'#endpoit-users'" class="menu-item">
                        /Users
                      </p>
                      <b-nav vertical class="w-100 left-menu">
                        <b-nav-item
                          ><p
                            v-scroll-to="'#endpoint-users-list'"
                            class="submenu-item"
                          >
                            <chevron-right-icon />List
                          </p></b-nav-item
                        >
                        <b-nav-item
                          ><p
                            v-scroll-to="'#endpoint-users-get'"
                            class="submenu-item"
                          >
                            <chevron-right-icon />Get
                          </p></b-nav-item
                        >
                        <b-nav-item
                          ><p
                            v-scroll-to="'#endpoint-users-create'"
                            class="submenu-item"
                          >
                            <chevron-right-icon />Create
                          </p></b-nav-item
                        >
                        <b-nav-item
                          ><p
                            v-scroll-to="'#endpoint-users-update'"
                            class="submenu-item"
                          >
                            <chevron-right-icon />Update
                          </p></b-nav-item
                        >
                        <b-nav-item
                          ><p
                            v-scroll-to="'#endpoint-users-delete'"
                            class="submenu-item"
                          >
                            <chevron-right-icon />Delete
                          </p></b-nav-item
                        >
                        <b-nav-item
                          ><p
                            v-scroll-to="'#endpoint-users-login'"
                            class="submenu-item"
                          >
                            <chevron-right-icon />Login
                          </p></b-nav-item
                        >
                        <b-nav-item
                          ><p
                            v-scroll-to="'#endpoint-users-logout'"
                            class="submenu-item"
                          >
                            <chevron-right-icon />Logout
                          </p></b-nav-item
                        >
                      </b-nav>
                    </b-nav-item>
                  </b-nav>
                </div>
                <div class="menu-left-footer">
                  <a href="https://b4ysolutions.com">
                    <img alt="Vue logo" src="../assets/b4y_small.png" class="logo-b4y" />
                  </a>
                </div>
              </div>
            </b-sidebar>
          </div>
          <!-- END - Left Menu -->
        </b-navbar-nav>
      </b-container>
    </b-navbar>
  </div>
</template>

<script>
import MenuIcon from "vue-material-design-icons/Menu.vue";
import ChevronRightIcon from "vue-material-design-icons/ChevronRight.vue";

export default {
  name: "Menu",
  components: {
    MenuIcon,
    ChevronRightIcon,
  },
};
</script>

<style  lang="scss">
nav.navbar {
  padding: 0px !important;
}
.menu-left {
  margin-top: -30px;
  height: calc(100% - 30px);
  overflow-y: auto;
}
.menu-left-footer {
  position: absolute;
  bottom: 0;
  text-align: center;
  width: 100%;
}
p.menu-item {
  background: #008b3a;
  color: #fff !important;
  padding: 8px 0 8px 10px;
  margin: 20px 0 0;
  font-size: 16px !important;
}
p.submenu-item {
  padding: 8px 0 4px 25px;
  margin: 0;
  border-right: 5px solid transparent;
}
p.submenu-item:hover {
  border-right: 5px solid #008b3a;
  color: #008b3a;
}
.logo-min {
  width: auto;
  height: 30px;
}
.logo-b4y {
  width: auto;
  height: 50px;
  margin: 5px 0;
}
.global-icon {
  font-size: 30px;
  padding: 1px;
  color: #008b3a;
  margin-right: 10px;
}
.navbar-expand .navbar-nav .nav-link {
  padding: 0px !important;
  margin: 0px !important;
}
a.nav-link {
  color: #252525 !important;
  text-align: left;
  padding: 0px;
}
a.nav-link:hover {
  color: #008b3a !important;
}
h1.page-title-left {
  color: #008b3a !important;
  font-size: 18px;
  margin: 0px 0px 0px 10px !important;
  cursor: pointer;

  img{
    margin-right: 8px;
  }
}
</style>
