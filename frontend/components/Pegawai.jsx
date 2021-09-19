import React, { Component } from 'react'
import styles from '../pages/index.module.css'
import Link from 'next/link'
const axios = require('axios').default;

export default class Pegawai extends Component {

  constructor(props) {
    super(props)
    this.state = {
      isLoad : false,
      items : [],
      id_pegawai : null
    }

    this.loadData = this.loadData.bind(this)
    this.idHandler = this.idHandler.bind(this)
  }

  loadData = (isLoad) => {
    if (!isLoad) {
      return <p>Loading...</p>
    } else {
      const data = this.state.items
      let show_data = (
        <h2 className="text-center">Kosong</h2>
      )
      if (data != null){
        show_data = data.map(data => 
          <div className="mt-2">
            <h2 className="text-sm">ID: {data.id}</h2>
            <h3>Nama: {data.nama}</h3>
            <p className="text-sm">Alamat: {data.alamat}</p>
            <p className="text-sm">Telepon: {data.telepon}</p>
          </div>
        )
      }
      return (
        <div className="grid grid-flow-row">
          <h1 className="text-center text-lg">Pegawai yang tersedia</h1>
          {show_data}
        </div>
      )
    }
  }
  
  idHandler = (event) => {
    let val = event.target.value
    if ((val !== "") || (val !== null)) {
      let link = "/api/pegawai?id="+val
      axios.get(link).then(res => {
        this.setState({items:res.data.data})
      }).catch(e => console.log(e))
    } else {
      let link = "/api/pegawai"
      axios.get(link).then(res => {
        this.setState({items:res.data.data})
      }).catch(e => console.log(e))
    }
  }


  componentDidMount() {
    axios.get(`/api/pegawai`).then(res => {
      this.setState({isLoad: true, items: res.data.data})
    }).catch(e => console.log(e))
  }

  render() {
    return (
      <div>
        <div className={styles.container}>
          <h1 className="text-center">
            Pegawai
          </h1>

          <p className="mt-5">
            Disini saya menggunakan /api/pegawai untuk melakukan fetching data-data
            pegawai.
          </p>

          <p className="mt-5">
            Anda bisa melakukan pencarian pegawai menggunakan ID
          </p>

          <div className="grid grid-flow-row w-1/2 ml-auto mr-auto mt-3">
            <label htmlFor="id_pegawai">
              <p className="text-center">ID Pegawai</p>
            </label>
            <input type="number" name="id_pegawai" id="id_pegawai"
              placeholder="Masukkan ID Pegawai" onChange={this.idHandler} />
            <Link href="/">
              <a href="/" className="text-center mt-3 text-blue-500">Kembali Ke Menu Utama</a>
            </Link>
          </div>

          
        </div>

        

        <div className={styles.container} style={{ marginTop: "1rem" }}>
          {this.loadData(this.state.isLoad)}
        </div>
      </div>
    )
  }
}
