import React from 'react';

class ImageList extends React.Component {
  constructor() {
    super();
    this.state = {
      images: []
    };
  };

  componentDidMount() {
    fetch('http://localhost:8080') // Hard-coded localhost just for MVP demo purposes
      .then(response => response.json())
      .then(data =>
        this.setState({
          images: data
        })
      );
  }

  handleClick(id) {
    console.log("hello", id)
    // TODO POST API TO UPDATE FLAGGED - ON SUCCESS, SETSTATE - ON FAIL, RENDER ERROR
  }


  render() {
    const hasImages = this.state.images.length > 0
    const imageList = this.state.images.map((i) => <li key={i.id}>
      <img src={i.url}></img>
      <a onClick={() => this.handleClick(i.id)} className="button is-primary"> Tag</a>
    </li>);

    if (hasImages) {
      return <div>{imageList}</div>
    } else {
      return <div>Oops, something went wrong! Please refresh!</div>
    }
  }
}

export default ImageList;