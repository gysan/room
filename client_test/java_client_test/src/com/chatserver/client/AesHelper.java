package com.chatserver.client;

import javax.crypto.Cipher;
import javax.crypto.spec.SecretKeySpec;
import javax.crypto.spec.IvParameterSpec;

//AES�ǶԳƼ����㷨
//����ʹ�� AES-128��key���ȣ�16, 24, 32 bytes ��Ӧ AES-128, AES-192, AES-256

public class AesHelper {
	private IvParameterSpec ivSpec;
	private SecretKeySpec keySpec;
	
	public AesHelper() {
		// Ĭ��key
		String key = "12345abcdef67890";
		try {
			this.keySpec = new SecretKeySpec(key.getBytes(), "AES");
			this.ivSpec = new IvParameterSpec(key.getBytes());
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	public AesHelper(String key) {
		try {
			this.keySpec = new SecretKeySpec(key.getBytes(), "AES");
			this.ivSpec = new IvParameterSpec(key.getBytes());
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	public byte[] encrypt(byte[] origData) {
		try {
			Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
			cipher.init(Cipher.ENCRYPT_MODE, this.keySpec, this.ivSpec);
			return cipher.doFinal(origData);
		} catch (Exception e) {
			e.printStackTrace();
		}
		return null;
	}

	public byte[] decrypt(byte[] crypted) {
		try {
			Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
			cipher.init(Cipher.DECRYPT_MODE, this.keySpec, this.ivSpec);
			return cipher.doFinal(crypted);
		} catch (Exception e) {
			e.printStackTrace();
		}
		return null;
	}
	
	public static void main(String[] args) {
		AesHelper aes1 = new AesHelper("12345abcdef67890");
		String data1 = "hello world�Ұ��㣬 ������ݣ�����һ������";
		byte[] encrypted1 = aes1.encrypt(data1.getBytes());
		byte[] decrypted1 = aes1.decrypt(encrypted1);
		System.out.println(new String(decrypted1));
		
		AesHelper aes2 = new AesHelper();
		String data2 = "hi, ���ǵڶ�������";
		byte[] encrypted2 = aes2.encrypt(data2.getBytes());
		byte[] decrypted2 = aes2.decrypt(encrypted2);
		System.out.println(new String(decrypted2));
	}
	
}
