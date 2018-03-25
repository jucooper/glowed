import React, { Component } from 'react';
import $ from 'jquery';
import './App.css';

class App extends Component {

  constructor(props){
    super();
    this.state = { player: null };
  }

  componentWillMount() {
    $.ajax({
      url: 'http://localhost:5000/',
      dataType: 'json',
      success: function(data) {
        this.setState({player: data.player});
      }.bind(this)
    });
  }

  render() {
    return (
      <div class="container">
        <div class="row">
          <div>
            {
              this.state.player != null
              ? 
              <div>
                <p>{this.state.player.displayName}</p>
                <p>{this.state.player.stats.wins}</p>
                <p>{this.state.player.stats.goals}</p>
                <p>{this.state.player.stats.mvps}</p>
                <p>{this.state.player.stats.saves}</p>
              </div>
              :
              <div>
                <p className="loader"></p>
              </div>
            }
            <div>
              <p>All content about Rocket League belongs to Psyonix, Inc.</p>
              <a href="https://rocketleaguestats.com/"><img src="./rls_partner_horizontal_small.png" alt="img"></img></a>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default App;
