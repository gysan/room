package gotest;

import com.chatserver.client.AesHelper;
import com.chatserver.client.ConvHelper;
import com.chatserver.client.PacketHelper;
import com.google.protobuf.InvalidProtocolBufferException;
import pb.*;

public class GoTest {
	public static void main(String[] args) {
		// ����client��������
		com.chatserver.client.TcpHelper client = new com.chatserver.client.TcpHelper();
		boolean re = client.connect("115.29.174.190", 8989);
		if (!re) {
			System.out.println("connect error");
			return;
		}
		
		// ��¼
		
		// ��ʼ����¼��
		Pb.PbClientLogin.Builder clientLoginBuilder = Pb.PbClientLogin.newBuilder();
		clientLoginBuilder.setUuid("hello, ����һ��java����");
		clientLoginBuilder.setVersion(3.14f);
		clientLoginBuilder.setTimestamp((int)(System.currentTimeMillis()/1000));
		
		// ���͵�¼��
		byte[] loginData = marshalPbClientLogin(clientLoginBuilder);
		re = client.send(loginData);
		if (!re) {
			System.out.println("send error");
			return;
		}
		
		// ��ȡ����
		byte[] readData = new byte[1024];
		int readSize = client.read(readData);
		if (-1 == readSize) {
			System.out.println("read error");
			return;
		}
		
		// ���յ�¼��ִ��
		// �ٶ���ȡ�����������ǵ�¼��ִ����ͬʱ�ٶ�readSize�պ��ǵ�¼��ִ���ĳ���
		byte[] len = new byte[4];
		byte[] type = new byte[4];
		byte[] data = new byte[readSize-8];
		System.arraycopy(readData, 0, len, 0, 4);
		System.arraycopy(readData, 4, type, 0, 4);
		System.arraycopy(readData, 8, data, 0, readSize-8);
		
		// len �� typeֱ�Ӷ�ȡ
		System.out.println("len== " + ConvHelper.bytesToInt(len));
		System.out.println("type== " + ConvHelper.bytesToInt(type));
		if (ConvHelper.bytesToInt(type) == PacketHelper.PK_ServerAcceptLogin) {
			System.out.println("this is PK_ServerAcceptLogin");
		}
		
		// ������¼��ִ��
		Pb.PbServerAcceptLogin serverAcceptLogin = unmarshalPbServerAcceptLogin(data);
		if (null == serverAcceptLogin) {
			System.out.println("unmarshalPbServerAcceptLogin error");
			return;
		}
		System.out.println(serverAcceptLogin.getLogin());
		System.out.println(serverAcceptLogin.getTipsMsg());
		System.out.println(serverAcceptLogin.getTimestamp());
		
		// �ر�����
		client.close();
	}
	
	// PbClientLogin 
	// ��protobuf�ṹ��builder �����л�����AES���ܣ� �ٴ��
	public static byte[] marshalPbClientLogin(Pb.PbClientLogin.Builder b) {
		AesHelper aes = new AesHelper();
		Pb.PbClientLogin info = b.build();
		byte[] encrypted = aes.encrypt(info.toByteArray());
		return new PacketHelper(PacketHelper.PK_ClientLogin, encrypted).getBytes();
	}
	
	// PbServerAcceptLogin
	// data�����������󴫽����ģ���data�Ƚ��ܣ��ٷ����л�Ϊprotobuf�ṹ
	public static Pb.PbServerAcceptLogin unmarshalPbServerAcceptLogin(byte[] data) {
		AesHelper aes = new AesHelper();
		byte[] decrypted = aes.decrypt(data);
		
		Pb.PbServerAcceptLogin rev = null;
		try {
			rev = Pb.PbServerAcceptLogin.parseFrom(decrypted);
		} catch (InvalidProtocolBufferException e) {
			return null;
		}
		return rev;
	}
}
