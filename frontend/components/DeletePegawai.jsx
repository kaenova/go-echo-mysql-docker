import axios from 'axios'
import React, { Component } from 'react'
import styles from '../pages/index.module.css'
import {motion} from 'framer-motion'


export default class DeletePegawai extends Component {

  constructor(props) {
    super(props)
    this.state = {
      berhasil_hapus: null,
      data: {
        nama: "lorem",
        alamat: "ipsum",
        telepon: "dayeuh kolot"
      },
      pop_confirmation : false
    }
  }


  berhasil_hapus = () => {
    if (this.state.berhasil_hapus == false) {
      return (
        <div className={styles.container} style={{ marginTop: "1rem" }}>
          <h2 className="text-red-500">Tidak berhasil Menghapus Pegawai</h2>
          <p>Harap periksa kembali nama dan nomor telepon pegawai</p>
        </div>
      )
    } else if (this.state.berhasil_hapus == true) {
      return (
        <div className={styles.container} style={{ marginTop: "1rem" }}>
          <h2 className="text-green-500">Berhasil Menghapus Pegawai</h2>
          <p>Data Pegawai:</p>
          <p>Nama:  {this.state.data.nama}</p>
          <p>Alamat:  {this.state.data.alamat}</p>
          <p>Telepon:  {this.state.data.telepon}</p>
        </div>
      )
    }
  }

  formHandler = (event) => {
    event.preventDefault()
    this.state.data = {
      nama : event.target.nama.value,
      telepon: event.target.telepon.value
    }
    this.setState({pop_confirmation : true})
  }


  hapus_pegawai = () => {
    let url = `/api/pegawai?nama=${this.state.data.nama}&telepon=${this.state.data.telepon}`
    axios.delete(url).then(res => {
      let data = res.data.data
      if (res.status == 200) {
        this.setState({
          berhasil_hapus: true,
          data : {
            nama: data.nama,
            alamat: data.alamat,
            telepon: data.telepon
          },
          pop_confirmation: false
        })
      } else {
        this.setState({
          berhasil_hapus : false,
          pop_confirmation: false
        })
      }
    }).catch(e => {
      this.setState({
        berhasil_hapus : false,
        pop_confirmation: false
      })
    })
  }

  confirmation = () => {
    if (this.state.pop_confirmation) {
      return (
        <div className="w-3/4 ml-auto mr-auto mt-5 grid grid-flow-row justify-center">
          <h2 className="text-red-600 text-sm text-center">Apakah anda yakin ingin menghapus pegawai tersebut?</h2> 
          <motion.button className="bg-red-500 hover:bg-red-600" onClick={this.hapus_pegawai}
          whileHover={{scale:1.1}}>Saya mengerti dan hapus</motion.button>
        </div>
      )
    }
  }

  

  render() {
    return (
      <div>
        <div className={styles.container}>
          <h1>Hapus Pegawai</h1>

          <form onSubmit={this.formHandler} className="mt-4 grid grid-flow-row">
            <label htmlFor="nama">
              Nama:
              <input type="text" name="nama" id="nama" placeholder="Masukkan nama lengkap" required/>
            </label>
            <label htmlFor="telepon">
              Telepon:
                <input type="number" name="telepon" id="telepon" placeholder="+628572830777" required/>
            </label>
            <motion.button type="submit" className="mt-5 w-1/2 mr-auto ml-auto"
            whileTap={{scale:0.9}}>Submit</motion.button>
          </form>

          {this.confirmation()}

        </div>

        {this.berhasil_hapus()}
      </div>
    )
  }
}