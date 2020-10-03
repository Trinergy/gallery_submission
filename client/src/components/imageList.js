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

  handleClick(id, flagged) {
    const formData = new FormData();
    formData.append('flagged', !flagged)

    fetch(`http://localhost:8080/image/${id}`, {
      method: 'PUT',
      body: formData
    })
    .then(response => {
      // This can be handled better with proper state management libarries - kept small for MVP
      var copy = [...this.state.images];
      var updatedImageIndex = copy.findIndex(image => image.id == id)
      copy[updatedImageIndex] = {...copy[updatedImageIndex], ...{flagged: !flagged}}

      this.setState({
        images: copy
      })
    })
    .catch(error => {
    });
  }


  render() {
    // Removed CSS management libraries to keep things simple
    const hasImages = this.state.images.length > 0
    const imageList = this.state.images.map( i => <li key={i.id}>
      <img border={i.flagged ? "10px solid" : ""} width="50%" height="auto" src={i.url}></img>
      <a onClick={() => this.handleClick(i.id, i.flagged)} className="button is-primary"> Tag</a>
    </li>);

    if (hasImages) {
      return <div>{imageList}</div>
    } else {
      return <div>Oops, something went wrong! Please refresh!</div>
    }
  }
}

export default ImageList;